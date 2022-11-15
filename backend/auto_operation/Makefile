test: staticcheck typecheck

typecheck:
	mypy main.py
staticcheck:
	pflake8 ./
format:
	isort .
	black .

