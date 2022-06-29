package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	if len(rhyme) == 0 {
		return nil
	}

	proverbs := make([]string, 0, len(rhyme))
	for i := 0; i < len(rhyme)-1; i++ {
		proverbs = append(proverbs, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
	}
	proverbs = append(proverbs, "And all for the want of a "+rhyme[0]+".")
	return proverbs
}
