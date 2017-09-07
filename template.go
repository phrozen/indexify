package main

const tmpl = `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Path | title}}</title>
	<style>
		html, body {
			width: 100%;
			margin: 10px;
			font-family: "Segoe UI", "Helvetica Neue", Arial, sans-serif;
			line-height: 1.05em;
			color: #333;
		}
		.left {
			text-align: left;
		}
		.right {
			text-align: right;
		}
		.center {
			text-align: center;
		}
		tbody > tr> td, th {
			padding: 5px 10px;
			line-height: 22px;
		}
		table tbody tr:nth-child(even) {
			background-color: #eee;
		}
		a {
			text-decoration: none;
			color: #339;
		}
		a:hover {
			text-decoration: underline;
		}
		/*
			Icons obtained from IcoMoon Free Pack -	https://icomoon.io
			CC BY 4.0 - https://creativecommons.org/licenses/by/4.0/
		*/
		i {
			display: block;
			width: 16px;
			height: 16px;
			background-position: center center;
			background-repeat: no-repeat;
			background-size: cover;
		}
		i.folder {
			background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAPElEQVQ4T2NkoBAwUqifAWTAfzyGLGBgYEjEZwkhA0B68RpCjAF4fTlqAAPBaCSYTEYDkUqBSDCk8SkAANzTDwbegmyWAAAAAElFTkSuQmCC');
		}
		i.file {
			background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAdElEQVQ4T2NkgIANDAwM/lA2LuoBAwNDIgMDwwFkBYxQzn8GBgYYG5cBIDUfGRgYChgYGBbAFJFqgCDUBRNghpBqAEi9AAMDA8g7ASDDSDHgAwMDAz+S/0CGKJJiAHrYgMNt1IDRMKBqOiAmO6OnxI2g/AAA+l0oESP+dzQAAAAASUVORK5CYII=');
		}
		i.text {
			background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAiElEQVQ4T+2T0Q1AMBCGv06ACTARK9jAKEYwAhsYhQ2YgBxtUpKGe5B40JfepX++3H/XMxynAwobh64RqIDBFxibrICLQwDRLEANtE6kBSS2gsZBtADRx4DYKQWmAcxA5PkTSK4BXHuz9+2bAPGWBmY5AZk/+lcs/BXcrKT3fPpIT9b5iu5lHzZPnC4R0QTp/AAAAABJRU5ErkJggg==');
		}
		i.image {
			background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAy0lEQVQ4T7WT4REBMRCFv6sAHVDCVWBUgBJ0oAOnAyqgA3SgAzqgBCpgHrsmMheT/JA/2bt5+Xb3ZVPxXntgbHFquwIz4BgKKvt4AB6nANLcgTmwdVEpoGcVrBxSCpC+C6idiWAlgBvQCfoTZFACiL15+VYCaIAFsAQU/wSoP/X6cdsOhOYnATp4AvrAKLj37ApcqGwyrjbXszxQ1kukPFslgqm6DTBNeaAxHbaMokZ9DewMIvNbPdDPnPU/QE5213y1kPOcY/hB7+EJ9808ERqIWksAAAAASUVORK5CYII=');
		}

		i.audio {
			background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAwklEQVQ4T5WTAQ0CMQxF3ykAHIACkIADkAAOcAISkAAKQAI4QAIogHzYkt7S0VuT5bbLz9tv13b84gSs0r72eQBb4GoFXTq8gbyvAaR5ATvgmEWtgElycMiQVoD0Y0DprAVrATyBkclPkFkLoKzNt24RYAnMgUVaU0B1UISA0rJb+H8OdMMduJl1Mc8dOvB6w/4bDNgA++RfT9hzXabgiXVTtXtLgCcWVJ2nsG3splATe/MR1iAYzn4fDBnnEnjWPHwATAI6ERNzzeQAAAAASUVORK5CYII=');
		}
		i.video {
			background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAoElEQVQ4T82T0Q3CMAxEXydoGYGJSkfoBh2FDWAE2IBRGKGdoNVBLIUokWK+4p8kkv109jkd33gAY7iXjjcwA684oQuPHbB7CaCcDViAuyV5Aaeg4GoQL0D5A6B2LoJ5ACvQR/0JcvYA0tl85tYmQNLSsOHdgCm2PtdCDqAiFcuBn5paQHH5agFSoO2TjX8psBkIogVSNGZjzXdOrX2qnQOUoDERdpKbEwAAAABJRU5ErkJggg==');
		}
		i.compressed {
			background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAAs0lEQVQ4T6WTURHCMBBEXxVQJKAAqgQs4AApSEBCcYAEcIAEUACzkGOOThJyQ3+apNvXvb1ex/sagXVal25XYAucvKBLmwfg1zq2vemluQM74GCHOYB/wYMEmCcHe4NEHUjfAypnI1jEwQ2YufoFWUQcTMN95RZx0AwQOdeFZkCpC82Avx0YQF88A0Ph1yyG6AG1LKqAWnd8Rtk2yvYyqS7AKlrCj6H8PP4qoWWcp+Cj5uEJd4U4EStFKFYAAAAASUVORK5CYII=');
		}
		i.github {
			background-image: url('data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAABJ0lEQVQ4T62TjU0CQRCFPyrQDsAK1AqUCrQDpQOpQK0A7AArUCtQKpAOgA6wAs1nZi6TFS8xcZJL5nZ23nvzswN+2iFwAVwC+toOeAZewu+yBk3+NTAriS28QFNgkYEK4OHVHkX7jrw7MZAAN8G8DakqOWgyP4LZ0oahZC7ACHgP2ffAXfgnwFuAnAOrqN/4bfhHAuSBdx8B2fusljoVQJazyDgNpj4AlalYWwrwWW63U/kNqMv5F4BNdFW2v5awVUFtiv0Y9zTAzXwCnMp30wXw5xV4CN/5z2NtVac5apfMCemnjbNpMh8Hu4q8ZLcrwLpRtpQwAZTmonjJFc3EmlOn5cZKsKtjE8QXlzvRjjQBZHadfVjdW6hM1umXjcqYZVpe9xINfAG+CUAPVSD+aQAAAABJRU5ErkJggg==');
		}
	</style>

</head>
<body>
	<h1>{{.Path | title}}</h1>
	<table id="list">
		<tbody>
			<tr>
				<th class="center"></th>
				<th class="left"><a href="#">Name</a></th>
				<th class="center"><a href="#">Size</a></th>
				<th class="right"><a href="#">Last Modified</a></th>
			</tr>
			{{range .List}}
				<tr>
					{{if .IsDir}}
						<td class="center"><i class="folder"></td>
						<td class="left"><a href="./{{.Name}}">{{.Name}}</a></td>
						<td></td>
					{{else}}
						<td class="center"><i class="{{.Name | type}}"></i></td>
						<td class="left"><a href="./{{.Name}}">{{.Name}}</a></td>
						<td class="center">{{.Size | formatSize}}</td>
					{{end}}
					<td class="right">{{.ModTime.Format "02 Jan 06 15:04 -0700"}}</td>
				</tr>
			{{end}}
		</tbody>
		<tfoot>
			<tr><td colspan="4" class="right">
				<span style="font-size: 10px;">Created with
				<a href="https://github.com/phrozen/indexify" target="_blank">indexify.</a>
				<i style="display: inline-block; width: 10px; height: 10px; margin: -1px;" class="github"></i>
				</span>
			</td></tr>
		</tfoot>
	</table>
</body>
</html>
`