package main

func readConfig() config {
	conf := config{
		path: "src",
		modules: []module{
			dictModule{
				typeId: identifier{
					moduleName: "Data.ByDate",
					name:       "ByDate",
				},
				publicKeyId: identifier{
					moduleName: "Calendar",
					name:       "Date",
				},
				privateKeyId: identifier{
					name: "String",
				},
			},
		},
	}
	return conf
}
