//go:build !solution

package permissions

func Grant(current, added Permission) Permission {
	return current | added
}

func Revoke(current, removed Permission) Permission {
	return current &^ removed
}

func Has(current, required Permission) bool {
	return current&required == required
}
