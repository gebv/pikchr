const vscode = require('vscode');
const path = require('path');
const httpRequest = require('request');

let previewPanel = null;
let serverURL = null;

/**
 * @param {vscode.ExtensionContext} context
 */
function activate(context) {
	let disposable = vscode.commands.registerCommand('pikchr.doPreviewSelectedTextDiagram', function () {
		serverURL = vscode.workspace.getConfiguration("pikchr").get("render_server")
		if (!serverURL) {
			vscode.window.showInformationMessage("Please to specify the URL of the render server");
			return;
		}

		vscode.window.showInformationMessage("Render server:", serverURL);

		// The code you place here will be executed every time your command is executed

		var editor = vscode.window.activeTextEditor;
		if (!editor) {
			vscode.window.showInformationMessage("No open text editor. Please to select the source code of the diagram");
			return;
		}

		var selection = editor.selection;
		var text = editor.document.getText(selection);
		if (!text) {
			vscode.window.showInformationMessage("Please to select the source code of the diagram");
			return;
		}

		if (!previewPanel) {
			previewPanel = vscode.window.createWebviewPanel(
				'pikchrDiagramPreviewer',
				'pikchr diagram preview',
				vscode.ViewColumn.Two,
				{
					enableScripts: true,
					retainContextWhenHidden: true,
					localResourceRoots: [vscode.Uri.file(path.join(context.extensionPath, 'media'))]
				}
			);
			previewPanel.onDidDispose(() => {
				previewPanel = null
			})
		}



		httpRequest.post(serverURL, {
			method: 'POST',
			body: {
				diagram_src: text
			},
			json: true // Autom
		}, (err,httpResponse,body) => {
			if (err) {
				vscode.window.showInformationMessage("Failed request:", err);
				return
			}
			if (httpResponse.statusCode != 200) {
				vscode.window.showInformationMessage("Failed render:", body.err);
				return
			}

			previewPanel.webview.html = webPageWith(body, previewPanel.webview, context)
		})
	});

	context.subscriptions.push(disposable);
}
exports.activate = activate;

// this method is called when your extension is deactivated
function deactivate() {
	previewPanel = null
}

module.exports = {
	activate,
	deactivate
}

function webPageWith(body, webview, context) {
	if (body.success) {
		const pageHTML = `<!DOCTYPE html>
		<html lang="en" style="background-color: white;">
		<head>
			<meta charset="UTF-8">

			<meta http-equiv="Content-Security-Policy" content="default-src 'none'; img-src 'self' data:;">

			<meta name="viewport" content="width=device-width, initial-scale=1.0">

			<title>Preview</title>
		</head>
		<body>
			`+body.img_inline_svg+`
		</body>
		</html>`
		return pageHTML
	}

	const onDiskPath = vscode.Uri.file(
        path.join(context.extensionPath, 'media', 'webviewpanel.css')
	  );

	const webviewpanelCSS = webview.asWebviewUri(onDiskPath)
	console.log('webviewpanelCSS', webviewpanelCSS)

	const pageHTML = `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">

		<meta http-equiv="Content-Security-Policy" content="default-src 'none'; img-src 'self' data:; style-src ${webview.cspSource};">

		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" type="text/css" href="${webviewpanelCSS}">

		<title>Error info</title>
	</head>
	<body>
		<pre>`+body.err+`</pre>
	</body>
	</html>`

	return pageHTML
}
