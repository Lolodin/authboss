<form action="{{mountpathed "login"}}" method="POST">
    {{with .error}}{{.}}<br />{{end}}
    <input type="text" class="form-control" name="email" placeholder="E-mail" value="{{.primaryIDValue}}"><br />
    <input type="password" class="form-control" name="password" placeholder="Password"><br />
	{{with .csrf_token}}<input type="hidden" name="csrf_token" value="{{.}}" />{{end}}
    {{with .modules}}{{with .remember}}<input visible="false" type="hidden" name="rm" value="true"></input><br />{{end}}{{end -}}
	{{with .redir}}<input type="hidden" name="redir" value="{{.}}" />{{end}}
    <button type="submit">Login</button>
    {{with .modules}}{{with .recover}}<br /><a href="{{mountpathed "recover"}}">Recover Account</a>{{end}}{{end -}}
</form>
