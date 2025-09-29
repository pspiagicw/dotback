package handler

// func Handle(ctx *kong.Context) {
//
// 	// if opts.ExampleConfig {
// 	// 	printExampleConfig()
// 	// } else if len(opts.Args) == 0 {
// 	// 	noCommandProvided(opts)
// 	// }
// 	//
// 	// executeCommand(opts)
// 	switch ctx.Command() {
// 	case "backup":
// 	case "config":
// 	}
// 	fmt.Println(ctx)
// }
//
// func executeCommand(opts *argparse.Opts) {
// 	cmd := opts.Args[0]
// 	opts.Args = opts.Args[1:]
//
// 	switch cmd {
// 	case "version":
// 		help.Version(opts.Version)
// 	case "backup":
// 		backup.Backup(opts)
// 	case "config":
// 		config.PrintConfig(opts)
// 	case "help":
// 		help.Handle(opts.Args, opts.Version)
// 	default:
// 		help.PrintHelp(opts.Version)
// 		goreland.LogFatal("No command named %s", cmd)
// 	}
// }
//
// func printExampleConfig() {
// 	help.HelpExampleConfig()
// 	os.Exit(0)
// }
//
// func noCommandProvided(opts *argparse.Opts) {
// 	help.PrintHelp(opts.Version)
// 	goreland.LogFatal("No command provided")
// }
