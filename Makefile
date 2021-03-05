wiki: SHELL:=/bin/bash
wiki:
	echo "Clean wiki"
	[ -d docs ]  && rm -fr ./docs

	echo "Clone Wiki"
	[ ! -d docs ]  && git clone git@github.com:josehbez/pm.wiki.git docs
	
	echo "Update Wiki files"
	sed -i 's/run().Execute()/doc.GenMarkdownTree(run(),".\/docs")/g' cmd/pm/main.go
	sed -i '10s/.*/\"github.com\/spf13\/cobra\/doc\"/' cmd/pm/main.go
	go run cmd/pm/main.go
	git checkout  cmd/pm/main.go	
	
	echo "Push commits"
	@if [[ `cd docs && git status --porcelain` ]]; then\
		cd docs ;\
		git add . ;\
		git commit -m "Update $(date)";\
		git push --set-upstream origin master;\
	fi
	