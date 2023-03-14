class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        # Adiciona todas as letras de s num mapa, como chaves;

        def count(st:str) -> dict:
            aux = {}
            for c in st:
                if c in aux:
                    aux[c] += 1
                else:
                    aux[c] = 1
            return aux
        

        count_s = count(s)
        count_t = count(t)

        return count_s == count_t
