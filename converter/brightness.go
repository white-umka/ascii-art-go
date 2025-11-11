package converter

func GetCharForBrightness(brightness float64) string { // 1 - черный, 100 - белый
    switch {
    case brightness > 90:
        return "@"  
    case brightness > 80:
        return "0"  
    case brightness > 65:
        return "O"
    case brightness > 45:
        return "o"
    case brightness > 30:
        return "*"
    case brightness > 15:
        return "."
    default:
        return " "  
    }
}