{{define "Theader"}}
<!doctype html>
<html lang="en">

	<head>
		<meta charset="UTF-8" />
		<title>CodeSystem</title>
		<link rel="stylesheet" type="text/css" href="../static/css/bootstrap.min.css" />
	</head>

	<body>
					<ul class="nav nav-tabs">
						<li class="navbar-header">
							<a href="#">CodeSystem</a>
						</li>
						
						<li {{if .IsHome}}class="active"{{end}}>
									<a href="/">首页</a>
						</li>
						 
						<li class="dropdown">
									<a href="/pro" data-toggle="dropdown">
										项目
										<span class="caret"></span>
									</a>
									<ul class="dropdown-menu">
										<li><a href="/proinfo">查询项目信息</a></li>
										<li class="divider"></li>
										<li><a href="/buildconf">生成配置文件</a></li>
									</ul>
						</li>
						<li {{if .ispro}}class="active"{{end}}><a href="/db">数据库</a></li>
						<li {{if .ispro}}class="active"{{end}}><a href="/monitor">监控</a></li>
					</ul>

{{end}}
