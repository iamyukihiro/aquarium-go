package logic

func InjectFiler() FilerInterface {
	return NewFiler()
}
