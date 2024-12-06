package label

import "fmt"

func directionToZPL(direction int) rune {
	switch direction {
	case 0:
		return 'N'
	case 90:
		return 'R'
	case 180:
		return 'I'
	case 270:
		return 'B'
	default:
		panic("Unsupported direction")

	}
}

func toTemplate(fieldID string) string {
	return fmt.Sprintf("{{ %s }}", fieldID)
}
