test: typecheck

typecheck:
	mypy main.py
format:
	isort --extend-skip-glob 'spec/**' .
	black --extend-exclude \.spec .

