package safe

func Go(call GoCall) {
	go func() {
		defer func() { _ = recover() }()
		call()
	}()
}

func GoWithDone(call GoCall, callback GoCallback) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				callback(err)
			} else {
				callback(nil)
			}
		}()
		call()
	}()
}
