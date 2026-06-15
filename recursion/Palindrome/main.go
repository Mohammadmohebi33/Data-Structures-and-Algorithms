package main

func isPalindrome(s string) bool {
	// شرط پایه: رشته خالی یا یک حرفی => پالیندرومه
	if len(s) <= 1 {
		return true
	}

	// شرط پایه: اگر اول و آخر متفاوت بود => نه پالیندروم
	if s[0] != s[len(s)-1] {
		return false
	}

	// مرحله بازگشتی: رشته رو بدون اول و آخر چک کن
	return isPalindrome(s[1 : len(s)-1])
}

func main() {

}
