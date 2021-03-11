GOBIN=$(shell go env GOPATH)/bin

go-install: SHELL:=/bin/bash
go-install: go-build
	@if [ -d "${GOBIN}" ] ; then\
		mv build/pm "${GOBIN}";\
		echo "instal on ${GOBIN}";\
		rm -r build;\
	fi

go-build: SHELL:=/bin/bash
go-build: 
	[ -d docs ]  && rm -r build
	[ ! -d build ] && mkdir build
	go build -o build/pm cmd/pm/main.go 
	chmod 755 build/pm

gh-wiki: SHELL:=/bin/bash
gh-wiki:
	#echo "Clean wiki"
	#[ -d docs ]  && rm -fr ./docs
	#
	#echo "Clone Wiki"
	#[ ! -d docs ]  && git clone git@github.com:josehbez/pm.wiki.git docs
	
	echo "Update Wiki files"
	sed -i 's/run().Execute()/doc.GenMarkdownTree(run(),".\/docs")/g' cmd/pm/main.go
	sed -i '6s/.*/\"github.com\/spf13\/cobra\/doc\"/' cmd/pm/main.go
	go run cmd/pm/main.go
	git checkout  cmd/pm/main.go	
	rm docs/Home.md 
	mv docs/pm.md docs/Home.md

	echo "Push commits"
	@if [[ `cd docs && git status --porcelain` ]]; then\
		cd docs ;\
		git add . ;\
		git commit -m "Update $(date)";\
		git push --set-upstream origin master;\
	fi
