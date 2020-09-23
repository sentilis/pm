package semver

//	sev init - inicializa semantic-versioning
//
//	sev -patch -minor -major
//
//	# release
//	sev -release -life-cycle custom|simplified
//
//	sev -release -stage  x stage, stable-release or stable-release-lts | alfa,beta,release-candidate,stable-release or stable-release-lts
//
//	sev -release -patch -minor -major
//
//	# build
//	sev -build -patch -minor -major
//	sev -build -- print build version
//
//	sev -- print version e.g
//		1.0.0
//		1.0.1-rc.1.0.1
//		1.0.1-rc.1+exp.sha.1324353
//
//	# :TODO
//	sev -git-tag -- create tag current branch last commit
//	sev -git-branch
//	# TODO
//	sev -changelog -- create CHANGELOG.rst file and add
//		lines
//
//		softname (version) (date September 15 2020) shagit
//		----------------
