#! /bin/bash
# Copyright (c) 2021 The Project Metadata Authors.

function gh_wiki (){

    echo "Clone Wiki"

	if [ -z "${GH_TOKEN}" ]; then 
	    [ ! -d docs ]  &&  git clone https://${GH_TOKEN}@github.com/josehbez/pm.wiki.git docs || echo "Exists docs"
    else 
        [ ! -d docs ]  &&  git clone git@github.com:josehbez/pm.wiki.git docs || echo "Exists docs"
	fi
	
	echo "Update cmd/pm/main.go"
	sed -i 's/run().Execute()/doc.GenMarkdownTree(run(),".\/docs")/g' cmd/pm/main.go
	sed -i '6s/.*/\"github.com\/spf13\/cobra\/doc\"/' cmd/pm/main.go
    
    echo "Autogen docs"
	go run cmd/pm/main.go
	git checkout  cmd/pm/main.go	
	rm docs/Home.md 
	mv docs/pm.md docs/Home.md
	
    echo "Markdow files remove end-links .md"
    find ./docs -type f  -name "*.md" | xargs sed -i 's/.md//g'

	echo "Push commits"
    cd docs
	if [[ `git status --porcelain` ]]; then
		git add .
		git commit -m "Update $(date)"
		git push --set-upstream origin master
	fi
    cd ..
}



if [ "$1" == "gh-wiki" ]; then 
    gh_wiki
fi 