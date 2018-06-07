ANTLR_BUILD_CONTAINER=antlrbuild

generate: ## generates antlr and visitor files from grammar
	docker build -t $(ANTLR_BUILD_CONTAINER) language
	cd language; docker run -t --rm -v `pwd`:/usr/src $(ANTLR_BUILD_CONTAINER) -Dlanguage=Go QueryLanguage.g4 -visitor -package language

