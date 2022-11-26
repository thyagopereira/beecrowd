# Desenhando Labirintos, Será possível que tem que implementar até o grafo ?

class Graph():

	def __init__(self):
		self.arestas = []
		self.nodes = []

	def add_node(self, node):
		self.nodes.append(node)

	def add_edge(self, begin, end):
		aresta =  str(begin) + "-" + str(end)
		self.arestas.append(aresta)
	
	def neighbors(self, n):
		neighbors = []

		for aresta in self.arestas:
			if str(n) in aresta:
				for elem in aresta.split("-"):
					if elem != n:
						neighbors.append(int(elem))

		neighbors.append(int(n))
		return neighbors

def tree(a):
	g = Graph()

	while(a > 0):
		bgn , end = [int (x) for x in input().split(" ")]

		if bgn not in g.nodes:
			g.add_node(bgn)
		if end not in g.nodes:
			g.add_node(end)

		g.add_edge(bgn, end)
		g.add_edge(end, bgn)
		a -= 1

	return g

def depth_search(graph, visited, n, result):

	visited[n] = True
	result.append(1)	
	neighbors = graph.neighbors(n)
	
	if len(neighbors) == 1 and visited[neighbors[0]]:  #Condição de parada é nó folha.
		return
	else:
		for neighbor in neighbors:
			if not visited[neighbor]:
				result.append(1)
				depth_search(graph, visited, neighbor, result)
	
def main():
	t = int(input())

	while (t > 0):

		n = int(input())
		v, a = [int (x) for x in input().split(" ")]

		graph = tree(a)

		visited =  [False] * v
		
		result = []
		depth_search(graph, visited, n, result)

		print(len(result) -1)
		t -=1 
	
main()