package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// panic("Please implement the WelcomeMessage() function")
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// panic("Please implement the AddBorder() function")
    return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    lines := strings.Split(oldMsg, "\n")
    var cleaned []string

    for _, line := range lines {
        trimmed := strings.TrimSpace(line)
        inner := strings.Trim(trimmed, "* ")
        if inner != "" {
            cleaned = append(cleaned, inner)
        }
    }

    return strings.Join(cleaned, "\n")
}

