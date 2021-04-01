package maptricks

// ReverseMap permite inverter um map[string]string, isto é, chave-valor é invertido para valor-chave. Use a função abaixo para reverter o mapa acima para chave: código e valor: nome. Pode ser útil para interpretar o retorno dos Correios.
func ReverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}
