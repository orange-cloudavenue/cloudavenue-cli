package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

type ModelEntry struct {
	Object        string `json:"object"`
	Type          string `json:"type,omitempty"`
	Documentation string `json:"documentation,omitempty"`
}

type Command struct {
	Namespace string       `json:"namespace,omitempty"`
	Resource  string       `json:"resource,omitempty"`
	Verb      string       `json:"verb,omitempty"`
	Model     []ModelEntry `json:"model,omitempty"`
}

type functionality struct {
	Title            string                   `json:"Title,omitempty"`
	Commands         map[string]Command       `json:"Commands,omitempty"`
	SubFunctionality map[string]functionality `json:"SubFunctionality,omitempty"`
}

var root map[string]functionality

const selectedFile = "../../selected_fields.json"

func main() {
	var dataPath string
	var port int
	flag.StringVar(&dataPath, "data", "functionalities.json", "path to functionalities.json")
	flag.IntVar(&port, "port", 8085, "http port")
	flag.Parse()

	if err := loadFunctionalities(dataPath); err != nil {
		log.Fatalf("load definitions: %v", err)
	}

	// Serve static web assets relative to the executable directory so the server
	// works regardless of the current working directory when the binary is run.
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/api/namespaces", namespacesHandler)
	http.HandleFunc("/api/commands", commandsHandler)
	http.HandleFunc("/api/model", modelHandler)
	http.HandleFunc("/api/save", handleSave)
	http.HandleFunc("/api/selection", selectionHandler)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("UI running on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// loadFunctionalities charge le fichier functionalities.json dans la variable root.
func loadFunctionalities(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	m := map[string]functionality{}
	if err := dec.Decode(&m); err != nil {
		return err
	}
	root = m
	return nil
}

// ListNamespaces retourne la liste des namespaces disponibles (triée)
func ListNamespaces() []string {
	out := []string{}
	for k := range root {
		out = append(out, k)
	}
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// ListCommands renvoie un slice de commandes (namespace, resource, verb)
func ListCommands(namespace string) ([]Command, error) {
	fn, ok := root[namespace]
	if !ok {
		return nil, fmt.Errorf("namespace %s not found", namespace)
	}
	out := []Command{}
	for v, cmd := range fn.Commands {
		if cmd.Verb == "" {
			cmd.Verb = v
		}
		out = append(out, cmd)
	}
	for resName, sub := range fn.SubFunctionality {
		for v, cmd := range sub.Commands {
			if cmd.Verb == "" {
				cmd.Verb = v
			}
			if cmd.Resource == "" {
				cmd.Resource = resName
			}
			out = append(out, cmd)
		}
	}
	// deterministic order
	sort.Slice(out, func(i, j int) bool {
		a := out[i]
		b := out[j]
		if a.Resource != b.Resource {
			return a.Resource < b.Resource
		}
		if a.Verb != b.Verb {
			return a.Verb < b.Verb
		}
		return a.Namespace < b.Namespace
	})
	return out, nil
}

// GetCommandModel retourne les ModelEntry pour namespace/resource/verb (resource peut être "")
func GetCommandModel(namespace, resource, verb string) ([]ModelEntry, error) {
	fn, ok := root[namespace]
	if !ok {
		return nil, fmt.Errorf("namespace %s not found", namespace)
	}

	// subfunctionality
	if resource != "" {
		if sub, ok := fn.SubFunctionality[resource]; ok {
			if cmd, ok2 := sub.Commands[verb]; ok2 {
				return cmd.Model, nil
			}
		}
	}

	// direct command
	if cmd, ok := fn.Commands[verb]; ok {
		return cmd.Model, nil
	}

	// fallback: search across subfunctionalities
	for _, sub := range fn.SubFunctionality {
		if cmd, ok := sub.Commands[verb]; ok {
			return cmd.Model, nil
		}
	}
	return nil, fmt.Errorf("command %s (resource=%s) not found in %s", verb, resource, namespace)
}

func namespacesHandler(w http.ResponseWriter, r *http.Request) {
	names := ListNamespaces()
	writeJSON(w, names)
}

func commandsHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("namespace")
	if ns == "" {
		writeJSON(w, map[string]string{"error": "missing namespace"})
		return
	}
	cmds, err := ListCommands(ns)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, cmds)
}

func modelHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("namespace")
	res := r.URL.Query().Get("resource")
	verb := r.URL.Query().Get("verb")
	if ns == "" || verb == "" {
		writeJSON(w, map[string]string{"error": "namespace and verb required"})
		return
	}
	model, err := GetCommandModel(ns, res, verb)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, model)
}

// FieldEntry représente un champ avec propriétés (compatible ancien/nouveau format)
type FieldEntry struct {
	Name       string            `json:"name"`
	ID         string            `json:"id,omitempty"`
	Display    string            `json:"display,omitempty"`
	Index      int               `json:"index,omitempty"`
	Expression string            `json:"expression,omitempty"`
	Bindings   map[string]string `json:"bindings,omitempty"`
	BuildedID  string            `json:"builded_id,omitempty"`
}

// FieldList accepte []string ou [{name,display}]
type FieldList []FieldEntry

func (fl *FieldList) UnmarshalJSON(data []byte) error {
	// essayer comme tableau générique
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var parsed FieldList
	for _, r := range raw {
		var s string
		if err := json.Unmarshal(r, &s); err == nil {
			parsed = append(parsed, FieldEntry{Name: s})
			continue
		}
		var obj FieldEntry
		if err := json.Unmarshal(r, &obj); err == nil {
			// accept object entries whether they include name or only builded_id
			if obj.Name != "" || obj.BuildedID != "" || obj.ID != "" {
				parsed = append(parsed, obj)
				continue
			}
		}
		// dernier recours : map
		var m map[string]interface{}
		if err := json.Unmarshal(r, &m); err == nil {
			idx := 0
			if rawIdx, ok := m["index"]; ok {
				switch v := rawIdx.(type) {
				case float64:
					idx = int(v)
				case int:
					idx = v
				case int32:
					idx = int(v)
				case int64:
					idx = int(v)
				}
			}
			if n, ok := m["name"].(string); ok {
				d := ""
				if dd, ok2 := m["display"].(string); ok2 {
					d = dd
				}
				id := ""
				if idVal, ok2 := m["id"].(string); ok2 {
					id = idVal
				}
				buildedID := ""
				if bid, ok2 := m["builded_id"].(string); ok2 {
					buildedID = bid
				}
				fe := FieldEntry{Name: n, Display: d, ID: id, BuildedID: buildedID, Index: idx}
				if ex, ok := m["expression"].(string); ok {
					fe.Expression = ex
				}
				if b, ok := m["bindings"].(map[string]interface{}); ok {
					fe.Bindings = map[string]string{}
					for k, v := range b {
						if s, ok2 := v.(string); ok2 {
							fe.Bindings[k] = s
						}
					}
				}
				parsed = append(parsed, fe)
				continue
			}
			// support object that contains only builded_id
			if bid, ok := m["builded_id"].(string); ok && bid != "" {
				fe := FieldEntry{BuildedID: bid, Index: idx}
				parsed = append(parsed, fe)
				continue
			}
			if id, ok := m["id"].(string); ok && id != "" {
				fe := FieldEntry{ID: id, Index: idx}
				parsed = append(parsed, fe)
				continue
			}
		}
		return fmt.Errorf("invalid field entry: %s", string(r))
	}
	*fl = parsed
	return nil
}

// SaveRequest attendu par /api/save
type SaveRequest struct {
	Namespace  string    `json:"namespace"`
	Resource   string    `json:"resource,omitempty"`
	Verb       string    `json:"verb"`
	Fields     FieldList `json:"fields"`
	WideFields FieldList `json:"wide_fields"`
	// Builded fields saved in simplified shape: id, display, expression, bindings (array)
	Builded []BuiltField `json:"builded_fields,omitempty"`
}

// BuiltField représente un champ composé enregistré, sans propriété "name" côté fichier
// Le client calcule un identifiant technique local si nécessaire.
type BuiltField struct {
	ID         string   `json:"id,omitempty"`
	Display    string   `json:"display,omitempty"`
	Expression string   `json:"expression,omitempty"`
	Bindings   []string `json:"bindings,omitempty"`
}

// Exemple de handler /api/save : ajuste selon votre code existant
func handleSave(w http.ResponseWriter, r *http.Request) {
	var req SaveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Construire la structure à écrire dans le fichier (conserve clefs historiques si besoin)
	outEntry := map[string]interface{}{
		"namespace": req.Namespace,
		"resource":  req.Resource,
		"verb":      req.Verb,
	}

	// Ici on écrit uniquement name/display pour champs simples, et id/display pour références vers des built-fields
	sel := make([]map[string]interface{}, 0, len(req.Fields))
	for i, f := range req.Fields {
		if f.Index <= 0 {
			f.Index = i + 1
		}
		m := map[string]interface{}{}
		if f.ID != "" || f.BuildedID != "" {
			// référence vers un built-field: normaliser sur la clé "id"
			if f.ID != "" {
				m["id"] = f.ID
			} else {
				m["id"] = f.BuildedID
			}
		} else {
			// champ normal du modèle
			m["name"] = f.Name
		}
		if f.Display != "" {
			m["display"] = f.Display
		}
		m["index"] = f.Index
		sel = append(sel, m)
	}
	wide := make([]map[string]interface{}, 0, len(req.WideFields))
	for i, f := range req.WideFields {
		if f.Index <= 0 {
			f.Index = i + 1
		}
		m := map[string]interface{}{}
		if f.ID != "" || f.BuildedID != "" {
			if f.ID != "" {
				m["id"] = f.ID
			} else {
				m["id"] = f.BuildedID
			}
		} else {
			m["name"] = f.Name
		}
		if f.Display != "" {
			m["display"] = f.Display
		}
		m["index"] = f.Index
		wide = append(wide, m)
	}

	// Écrire les objets enrichis
	outEntry["selected_fields"] = sel
	outEntry["selected_fields_wide"] = wide
	// built/composed fields created via the UI (simplified schema)
	if len(req.Builded) > 0 {
		bf := make([]map[string]interface{}, 0, len(req.Builded))
		for _, f := range req.Builded {
			m := map[string]interface{}{}
			if f.ID != "" {
				m["id"] = f.ID
			}
			if f.Display != "" {
				m["display"] = f.Display
			}
			if f.Expression != "" {
				m["expression"] = f.Expression
			}
			if len(f.Bindings) > 0 {
				m["bindings"] = f.Bindings
			} else {
				// ensure bindings exists as [] if not provided
				m["bindings"] = []string{}
			}
			bf = append(bf, m)
		}
		outEntry["builded_fields"] = bf
	}

	// legacy fields removed: we only store enriched objects (name + display)

	// Lecture existant file (si vous stockez un tableau Commands)
	var root map[string]interface{}
	if b, err := ioutil.ReadFile(selectedFile); err == nil {
		_ = json.Unmarshal(b, &root)
	} else {
		root = map[string]interface{}{"Commands": []interface{}{}}
	}

	// remplacer ou ajouter l'entrée correspondante
	cmds, _ := root["Commands"].([]interface{})
	replaced := false
	for i, c := range cmds {
		if m, ok := c.(map[string]interface{}); ok {
			if m["namespace"] == req.Namespace && m["resource"] == req.Resource && m["verb"] == req.Verb {
				cmds[i] = outEntry
				replaced = true
				break
			}
		}
	}
	if !replaced {
		cmds = append(cmds, outEntry)
	}
	root["Commands"] = cmds

	// écrire fichier en pretty json
	outB, _ := json.MarshalIndent(root, "", "  ")
	if err := ioutil.WriteFile(selectedFile, outB, 0644); err != nil {
		http.Error(w, "failed to write file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// répondre
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"file": selectedFile})
}

// selectionHandler retourne la sélection enregistrée pour une commande donnée (both lists)
func selectionHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("namespace")
	res := r.URL.Query().Get("resource")
	verb := r.URL.Query().Get("verb")
	if ns == "" || verb == "" {
		writeJSON(w, map[string]string{"error": "namespace and verb required"})
		return
	}
	sel, wide, built, err := getSelection(ns, res, verb)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, map[string]interface{}{"fields": sel, "fields_wide": wide, "builded_fields": built})
}

// getSelection lit le fichier selected_fields.json et retourne les champs sélectionnés pour une commande donnée.
func getSelection(namespace, resource, verb string) (fields []FieldEntry, wideFields []FieldEntry, builded []BuiltField, err error) {
	path := selectedFile
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("cannot read %s: %w", path, err)
	}
	var root struct {
		Commands []map[string]interface{} `json:"Commands"`
	}
	if err := json.Unmarshal(b, &root); err != nil {
		return nil, nil, nil, fmt.Errorf("invalid json: %w", err)
	}
	for _, c := range root.Commands {
		ns, _ := c["namespace"].(string)
		res, _ := c["resource"].(string)
		v, _ := c["verb"].(string)
		if ns == namespace && res == resource && v == verb {
			// champs principaux
			fields = parseFieldListFromInterface(c["selected_fields"])
			wideFields = parseFieldListFromInterface(c["selected_fields_wide"])
			builded = parseBuiltFieldsFromInterface(c["builded_fields"])
			return fields, wideFields, builded, nil
		}
	}
	return nil, nil, nil, fmt.Errorf("no selection found for %s/%s/%s", namespace, resource, verb)
}

// parseFieldListFromInterface convertit une interface{} (array) en []FieldEntry
func parseFieldListFromInterface(val interface{}) []FieldEntry {
	arr, ok := val.([]interface{})
	if !ok {
		return nil
	}
	out := make([]FieldEntry, 0, len(arr))
	for _, item := range arr {
		m, ok := item.(map[string]interface{})
		if ok {
			name, _ := m["name"].(string)
			display, _ := m["display"].(string)
			id, _ := m["id"].(string)
			buildedID, _ := m["builded_id"].(string)
			idx := 0
			if rawIdx, ok := m["index"]; ok {
				switch v := rawIdx.(type) {
				case float64:
					idx = int(v)
				case int:
					idx = v
				case int32:
					idx = int(v)
				case int64:
					idx = int(v)
				}
			}
			// normalize: if "id" is present, treat it as a built reference when builded_id is absent
			if id != "" && buildedID == "" {
				buildedID = id
			}
			fe := FieldEntry{Name: name, ID: id, Display: display, BuildedID: buildedID, Index: idx}
			out = append(out, fe)
		}
	}
	if len(out) == 0 {
		return out
	}
	sort.SliceStable(out, func(i, j int) bool {
		ii := out[i].Index
		ji := out[j].Index
		if ii == ji {
			return i < j
		}
		if ii == 0 {
			return false
		}
		if ji == 0 {
			return true
		}
		return ii < ji
	})
	used := make(map[int]bool, len(out))
	for _, fe := range out {
		if fe.Index > 0 {
			used[fe.Index] = true
		}
	}
	nextIdx := 1
	for i := range out {
		if out[i].Index == 0 {
			for used[nextIdx] {
				nextIdx++
			}
			out[i].Index = nextIdx
			used[nextIdx] = true
			nextIdx++
		} else if out[i].Index >= nextIdx {
			nextIdx = out[i].Index + 1
		}
	}
	return out
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

// parseBuiltFieldsFromInterface convertit une interface{} (array) en []BuiltField
func parseBuiltFieldsFromInterface(val interface{}) []BuiltField {
	arr, ok := val.([]interface{})
	if !ok {
		return nil
	}
	out := make([]BuiltField, 0, len(arr))
	for _, item := range arr {
		m, ok := item.(map[string]interface{})
		if ok {
			var bf BuiltField
			if id, _ := m["id"].(string); id != "" {
				bf.ID = id
			}
			if dsp, _ := m["display"].(string); dsp != "" {
				bf.Display = dsp
			}
			if ex, _ := m["expression"].(string); ex != "" {
				bf.Expression = ex
			}
			// bindings may be an array of strings
			if bArr, ok2 := m["bindings"].([]interface{}); ok2 {
				tmp := make([]string, 0, len(bArr))
				for _, v := range bArr {
					if s, ok3 := v.(string); ok3 {
						tmp = append(tmp, s)
					}
				}
				bf.Bindings = tmp
			}
			out = append(out, bf)
		}
	}
	return out
}
