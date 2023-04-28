package q4

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.


func ClassifyPrices(prices []int) (int, error) {
	n := len(prices)
	if n == 0 {
		return 0, fmt.Errorf("true")
	} else if n == 1 {
		return 3, nil
	}
	primeiro := 0
	if prices[0] < prices[1] {
		primeiro = 1
	} else if prices[0] > prices[1] {
		primeiro = 2
	}
	for i := 2; i < n; i++ {
		if prices[i-1] > prices[i] {
			if primeiro == 1 {
				return 3, nil
			}
			return 2, nil
		}
		if prices[i-1] < prices[i] {
			if primeiro == 2 {
				return 3, nil
			}
			return 1, nil
		}
	}
	if primeiro == 1 {
		return 1, nil
	} else {
		return 2, nil
	}
	return 0, nil
}

