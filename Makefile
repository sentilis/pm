wiki: SHELL:=/bin/bash
wiki:
	sed -i 's/run().Execute()/doc.GenMarkdownTree(run(),".\/docs")/g' cmd/pm/main.go
	sed -i '10s/.*/\"github.com\/spf13\/cobra\/doc\"/' cmd/pm/main.go
	go run cmd/pm/main.go
	git checkout  cmd/pm/main.go	
	cd docs
	@if [[ `git status --porcelain` ]]; then\
		git add .;\
    	git commit -m "Update Wiki $(date)";\
    	git push --set-upstream origin master;\
	fi
	cd ..