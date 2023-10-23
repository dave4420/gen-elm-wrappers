package main

func readConfig() config {
	conf := config{
		path: "src",
		modules: []module{
			dictModule{
				wrapperType: identifier{
					moduleName: "Data.ByDate",
					name:       "ByDate",
				},
				publicKeyType: identifier{
					moduleName: "Calendar",
					name:       "Date",
				},
				privateKeyType: identifier{
					name: "String",
				},
				wrapKeyFn: identifier{
					moduleName: "Data.Date",
					name:       "maybeDateFromIsoString",
				},
				unwrapKeyFn: identifier{
					moduleName: "Data.Date",
					name:       "isoStringFromDate",
				},
			},
		},
	}
	return conf
}
