from common import read_input, solve

prefs = read_input()
prefs['me']  # For the side-effect on defaultdict

print solve(prefs)
