package shortner

func NewURLShortnerModule() URLShortner {
	store := NewURLShortnerStore()
	service := NewShortnerService(store)

	return NewURLShortner(service)
}
