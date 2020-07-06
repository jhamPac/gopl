package main

var listHTML = template.Must(template.New("simcart").Parse(`
<html>
<body>
<table>
	<tr>
		<th>item</th>
		<th>price</th>
	</tr>
{{range $k, $v := .}}
	<tr>
		<td>{{$k}}</td>
		<td>{{$v}}</td>
	</tr>
{{end}}
</table>
</body>
</html>
`))