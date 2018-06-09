package nidc

import (
	"fmt"
	"strings"
)

type Table struct {
	*Identity
}

func (table *Table) returnSelect() string {
	alias := table.Alias
	if alias == "" {
		alias = table.Name
	}

	if table.Field == nil || len(table.Field) == 0 {
		return fmt.Sprintf(`%s.*`, alias)
	}

	fields := make([]string, len(table.Field))
	for i, field := range table.Field {
		fields[i] = fmt.Sprintf(`%s.%s`, alias, field)
	}
	return strings.Join(fields, `, `)
}

func (table *Table) returnFrom() string {
	if table.Alias != "" {
		return fmt.Sprintf(`%s AS %s`, table.Name, table.Alias)
	}
	return fmt.Sprintf(`%s `, table.Name)
}

func (table *Table) returnWhere() string {
	if len(table.Filter) == 0 {
		return ""
	}

	alias := table.Alias
	if alias == "" {
		alias = table.Name
	}

	filters := make([]string, len(table.Filter))
	for i, filter := range table.Filter {
		filters[i] = fmt.Sprintf(`%s.%s = %s`, alias, filter.Name, filter.Value)
	}

	return strings.Join(filters, ` AND `)
}

