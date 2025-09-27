package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/orange-cloudavenue/cloudavenue-sdk-go-v2/commands"
	"github.com/urfave/cli/v3"
)

var _ output = (*outputTemplate)(nil)

type outputTemplate struct{}

func OutputTemplate() output {
	return &outputTemplate{}
}

type (
	outputTemplateField struct {
		Name  string
		Value string
	}

	outputTemplateFields []outputTemplateField

	outputTemplateData struct {
		Fields outputTemplateFields
		MaxLen int
	}
)

// Modifié pour utiliser le padding dans le template Go
var outputTemplateTmpl = `
{{- range .Fields }}
{{- if .Name }}
	{{- if .Value }}{{ pad .Name $.MaxLen }}  {{ else }}{{ .Name }}{{ end }}
{{- end }}{{ printf "%s\n" .Value }}
{{- end }}
`

func (o *outputTemplate) Print(f *field, result any, cmd *cli.Command, cavcmd commands.Command) {
	fields := outputTemplateFields{}
	maxLen := 0
	orderedSelected := make([]fieldDetail, len(f.SelectedFields))
	copy(orderedSelected, f.SelectedFields)
	sort.SliceStable(orderedSelected, func(i, j int) bool {
		a := orderedSelected[i].Index
		b := orderedSelected[j].Index
		aValid := a > 0
		bValid := b > 0
		switch {
		case aValid && bValid:
			if a != b {
				return a < b
			}
			return false
		case aValid && !bValid:
			return true
		case !aValid && bValid:
			return false
		default:
			return false
		}
	})
	// Regroupement pour les champs indexés
	type idxGroup struct {
		orderSuffix []string                  // ordre des suffixes (ex: ["ID", "Name"])
		values      map[int]map[string]string // index -> (suffix -> value)
	}
	groups := map[string]*idxGroup{}
	groupOrder := []string{}
	// Groupes pour objets hiérarchiques (ex: compute_capacity.cpu.limit)
	type objSubGroup struct {
		leafOrder []string
		leaves    map[string]string
	}
	type objGroup struct {
		subOrder      []string
		subs          map[string]*objSubGroup
		rootLeafOrder []string
		rootLeaves    map[string]string
	}
	objGroups := map[string]*objGroup{}
	objGroupOrder := []string{}
	// Humanize utilitaire: compute_capacity -> "Compute Capacity", cpu -> "CPU"
	humanize := func(s string) string {
		s = strings.ReplaceAll(s, "_", " ")
		parts := strings.Fields(s)
		for i, p := range parts {
			up := strings.ToUpper(p)
			if len(p) <= 3 || p == up {
				parts[i] = up
			} else {
				parts[i] = strings.ToUpper(p[:1]) + strings.ToLower(p[1:])
			}
		}
		return strings.Join(parts, " ")
	}

	// 2. Remplir les champs sans padding (le padding sera fait dans le template)
	for _, fi := range orderedSelected {
		// Builded field
		if fi.ID != "" {
			for _, bf := range f.BuildedFields {
				if bf.ID == fi.ID {
					expr := bf.Expression
					for _, key := range bf.Bindings {
						v, err := commands.GetValueAtPath(result, key)
						if err != nil {
							v = key
						}
						expr = strings.ReplaceAll(expr, "${"+key+"}", fmt.Sprintf("%v", v))
					}
					name := bf.Display
					fields = append(fields, outputTemplateField{
						Name:  name,
						Value: expr,
					})
					break
				}
			}
			continue
		}

		if strings.Contains(fi.Name, "{index}") {
			// Tenter de regrouper par préfixe de display
			display := fi.Display
			grpKey := ""
			suffix := ""
			if sp := strings.LastIndex(display, "."); sp > 0 {
				grpKey = display[:sp]
				suffix = strings.TrimSpace(display[sp+1:])
			} else {
				// Pas de séparation trouvée: fallback ancien rendu
				values, err := commands.GetAllValuesAtTarget(result, fi.Name)
				if err != nil {
					continue
				}
				for i, v := range values {
					fields = append(fields, outputTemplateField{
						Name:  fmt.Sprintf("%s[%d]", fi.Display, i),
						Value: fmt.Sprintf("%v", v),
					})
				}
				continue
			}

			values, err := commands.GetAllValuesAtTarget(result, fi.Name)
			if err != nil {
				continue
			}

			g, ok := groups[grpKey]
			if !ok {
				g = &idxGroup{orderSuffix: []string{}, values: map[int]map[string]string{}}
				groups[grpKey] = g
				groupOrder = append(groupOrder, grpKey)
			}

			// Assurer l'ordre des suffixes
			present := false
			for _, s := range g.orderSuffix {
				if s == suffix {
					present = true
					break
				}
			}
			if !present {
				g.orderSuffix = append(g.orderSuffix, suffix)
			}

			for i, v := range values {
				if _, ok := g.values[i]; !ok {
					g.values[i] = map[string]string{}
				}
				g.values[i][suffix] = fmt.Sprintf("%v", v)
			}
			continue
		}

		// Groupement hiérarchique pour objets (ex: compute_capacity.cpu.Limit)
		if strings.Contains(fi.Name, ".") && !strings.Contains(fi.Name, "{index}") {
			segs := strings.Split(fi.Name, ".")
			if len(segs) >= 2 {
				parents := segs[:len(segs)-1]
				leaf := segs[len(segs)-1]

				// Top-level group label
				top := humanize(parents[0])
				// Sub-key for remaining parents (joined) or empty if none
				var subKey string
				if len(parents) > 1 {
					rest := make([]string, 0, len(parents)-1)
					for _, p := range parents[1:] {
						rest = append(rest, humanize(p))
					}
					subKey = strings.Join(rest, " ")
				} else {
					subKey = ""
				}

				leafLabel := strings.TrimSpace(fi.Display)
				if leafLabel == "" {
					leafLabel = humanize(leaf)
				}

				v, err := commands.GetValueAtPath(result, fi.Name)
				if err != nil {
					continue
				}

				// Ensure group exists
				g, ok := objGroups[top]
				if !ok {
					g = &objGroup{subOrder: []string{}, subs: map[string]*objSubGroup{}, rootLeafOrder: []string{}, rootLeaves: map[string]string{}}
					objGroups[top] = g
					objGroupOrder = append(objGroupOrder, top)
				}

				if subKey == "" {
					if _, exists := g.rootLeaves[leafLabel]; !exists {
						g.rootLeafOrder = append(g.rootLeafOrder, leafLabel)
					}
					g.rootLeaves[leafLabel] = fmt.Sprintf("%v", v)
				} else {
					sg, ok := g.subs[subKey]
					if !ok {
						sg = &objSubGroup{leafOrder: []string{}, leaves: map[string]string{}}
						g.subs[subKey] = sg
						// Maintain subOrder only once per new subKey
						present := false
						for _, sk := range g.subOrder {
							if sk == subKey {
								present = true
								break
							}
						}
						if !present {
							g.subOrder = append(g.subOrder, subKey)
						}
					}
					if _, exists := sg.leaves[leafLabel]; !exists {
						sg.leafOrder = append(sg.leafOrder, leafLabel)
					}
					sg.leaves[leafLabel] = fmt.Sprintf("%v", v)
				}
				continue
			}
		}

		// Simple field
		v, err := commands.GetValueAtPath(result, fi.Name)
		if err != nil {
			continue
		}
		fields = append(fields, outputTemplateField{
			Name: func() string {
				if fi.Display != "" {
					return fi.Display
				}
				return humanize(fi.Name)
			}(),
			Value: fmt.Sprintf("%v", v),
		})
	}

	// Émettre les groupes indexés à la fin, strictement selon l'ordre de première apparition
	for _, grpKey := range groupOrder {
		g := groups[grpKey]
		// Index triés
		idxs := make([]int, 0, len(g.values))
		for i := range g.values {
			idxs = append(idxs, i)
		}
		sort.Ints(idxs)
		// header with count
		fields = append(fields, outputTemplateField{Name: fmt.Sprintf("%s (%d)", grpKey, len(idxs)), Value: ""})
		for _, i := range idxs {
			row := g.values[i]
			// utiliser l'ordre de suffixes enregistré
			for _, suf := range g.orderSuffix {
				if val, ok := row[suf]; ok {
					fields = append(fields, outputTemplateField{
						Name:  fmt.Sprintf("  [%d] %s", i, suf),
						Value: val,
					})
				}
			}
		}
	}

	// Dérouler les groupes hiérarchiques d'objets (compute_capacity, properties, ...)
	for _, top := range objGroupOrder {
		g := objGroups[top]
		// Top-level header
		fields = append(fields, outputTemplateField{Name: top, Value: ""})
		// Root leaves (top-level properties directly under top)
		for _, leaf := range g.rootLeafOrder {
			fields = append(fields, outputTemplateField{Name: strings.Repeat(" ", 2) + leaf, Value: g.rootLeaves[leaf]})
		}
		// Sub-groups (like CPU, Memory)
		for _, sub := range g.subOrder {
			fields = append(fields, outputTemplateField{Name: strings.Repeat(" ", 2) + sub, Value: ""})
			sg := g.subs[sub]
			for _, leaf := range sg.leafOrder {
				fields = append(fields, outputTemplateField{Name: strings.Repeat(" ", 4) + leaf, Value: sg.leaves[leaf]})
			}
		}
	}

	// Insérer une ligne vide entre sections de premier niveau (headers non indentés)
	if len(fields) > 0 {
		withBlanks := make([]outputTemplateField, 0, len(fields)+6)
		prevWasBlank := false
		for i, fld := range fields {
			isTopHeader := fld.Value == "" && fld.Name != "" && (len(fld.Name) == len(strings.TrimLeft(fld.Name, " ")))
			if isTopHeader && i != 0 && !prevWasBlank {
				withBlanks = append(withBlanks, outputTemplateField{Name: "", Value: ""})
				prevWasBlank = true
			}
			withBlanks = append(withBlanks, fld)
			prevWasBlank = false
		}
		fields = withBlanks
	}

	// Réordonner: s'assurer que "Storage profile" (champ simple) vient après le bloc "Networks"
	{
		spIndex := -1
		netStart := -1
		netEnd := -1
		for i, fld := range fields {
			if spIndex == -1 && strings.EqualFold(strings.TrimSpace(fld.Name), "Storage profile") && fld.Value != "" {
				spIndex = i
			}
			if netStart == -1 && strings.EqualFold(strings.TrimSpace(fld.Name), "Networks") && fld.Value == "" {
				netStart = i
			}
		}
		if spIndex != -1 && netStart != -1 && spIndex < netStart {
			// calculer fin du bloc Networks
			netEnd = netStart + 1
			for netEnd < len(fields) {
				if fields[netEnd].Value == "" && netEnd != netStart {
					// prochain header => fin du bloc
					break
				}
				netEnd++
			}
			// extraire Storage profile et l'insérer après le bloc Networks
			sp := fields[spIndex]
			fields = append(fields[:spIndex], fields[spIndex+1:]...)
			if netEnd > spIndex {
				// ajuster car on a retiré un élément avant netEnd
				netEnd--
			}
			// insérer à netEnd
			if netEnd >= len(fields) {
				fields = append(fields, sp)
			} else {
				fields = append(fields[:netEnd], append([]outputTemplateField{sp}, fields[netEnd:]...)...)
			}
		}
	}

	funcMap := template.FuncMap{
		"pad": func(s string, length int) string {
			return fmt.Sprintf("%-*s", length, s)
		},
	}

	tmpl, err := template.New("output").Funcs(funcMap).Parse(outputTemplateTmpl)
	if err != nil {
		return
	}

	// Recalculer maxLen d'après les champs finaux, uniquement pour les lignes avec valeur
	for _, fi := range fields {
		if fi.Value == "" {
			continue
		}
		if len(fi.Name) > maxLen {
			maxLen = len(fi.Name)
		}
	}

	data := outputTemplateData{
		Fields: fields,
		MaxLen: maxLen,
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		return
	}
}
