package main

func readConfig() config {
	conf := config{
		path: "src",
		modules: []module{
			dictModule{
				wrapperType: identifier{
					moduleName: "Type.DictInt",
					name:       "DictInt",
				},
				publicKeyType: identifier{
					name: "Int",
				},
				privateKeyType: identifier{
					name: "String",
				},
				wrapKeyFn: identifier{
					moduleName: "String",
					name:       "toInt",
				},
				unwrapKeyFn: identifier{
					moduleName: "String",
					name:       "fromInt",
				},
			},
		},
	}
	return conf
}
