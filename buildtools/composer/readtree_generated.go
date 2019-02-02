// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package composer

// level is 1-indexed
type LineParser func(line string) (level int, node Package, err error)

func ReadPackageTree(lines []string, parser LineParser) ([]Package, map[Package][]Package, error) {
	var imports []Package
	edges := make(map[Package]map[Package]bool)
	parents := []Package{}

	for _, line := range lines {
		level, node, err := parser(line)
		if err != nil {
			return nil, nil, err
		}

		// Add to graph.
		if len(parents) >= level {
			parents = parents[:level-1]
		}
		if level == 1 {
			imports = append(imports, node)
		} else {
			parent := parents[len(parents)-1]
			_, ok := edges[parent]
			if !ok {
				edges[parent] = make(map[Package]bool)
			}
			edges[parent][node] = true
		}
		parents = append(parents, node)
	}

	graph := make(map[Package][]Package)
	for parent, children := range edges {
		for child := range children {
			graph[parent] = append(graph[parent], child)
			_, ok := graph[child]
			if !ok {
				graph[child] = nil
			}
		}
	}
	for _, i := range imports {
		_, ok := edges[i]
		if !ok {
			graph[i] = nil
		}
	}

	return imports, graph, nil
}
