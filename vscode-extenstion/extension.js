// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
const vscode = require('vscode');
const httpRequest = require('request');

let previewPanel = null;
let serverURL = null;

// this method is called when your extension is activated
// your extension is activated the very first time the command is executed

/**
 * @param {vscode.ExtensionContext} context
 */
function activate(context) {

	// console.log("pikchr", pikchr)

	// Use the console to output diagnostic information (console.log) and errors (console.error)
	// This line of code will only be executed once when your extension is activated
	console.log('Congratulations, your extension "pikchr" is now active!');

	// The command has been defined in the package.json file
	// Now provide the implementation of the command with  registerCommand
	// The commandId parameter must match the command field in package.json
	let disposable = vscode.commands.registerCommand('pikchr.doPreviewSelectedTextDiagram', function () {
		// serverURL = vscode.workspace.getConfiguration("pikchr").get("server_url")
		// if (!serverURL) {
		// 	vscode.window.showInformationMessage("Please to specify the URL of the server");
		// 	return;
		// }

		// The code you place here will be executed every time your command is executed

		var editor = vscode.window.activeTextEditor;
		if (!editor) {
			vscode.window.showInformationMessage("No open text editor");
			return;
		}

		var selection = editor.selection;
		var text = editor.document.getText(selection);
		if (!text) {
			vscode.window.showInformationMessage("Please to select the source code of the diagram");
			return;
		}

		// Display a message box to the user
		// vscode.window.showInformationMessage("");

		httpRequest.post('https://tgentrypoint-dot-fader4.ew.r.appspot.com', {
			method: 'POST',
			body: {
				some: 'payload'
			},
			json: true // Autom
		}, (err,httpResponse,body) => {
			console.log('err', err)
			console.log('httpResponse', httpResponse)
			console.log('body', body)
		})

		if (!previewPanel) {
			previewPanel = vscode.window.createWebviewPanel(
				'pikchrDiagramPreviewer',
				'pikchr diagram preview',
				vscode.ViewColumn.Two,
				{
					enableScripts: true,
					retainContextWhenHidden: true,
				}
			);
			previewPanel.onDidDispose(() => {
				previewPanel = null
			})
		}

		const pageHTML = `<!DOCTYPE html>
		<html lang="en" style="background-color: white;">
		<head>
			<meta charset="UTF-8">

			<meta http-equiv="Content-Security-Policy" content="default-src 'none'; img-src 'self' data:;">

			<meta name="viewport" content="width=device-width, initial-scale=1.0">

			<title>Cat Coding</title>
		</head>
		<body>
			<img src="`+svgb64+`"></img>
		</body>
		</html>`

		previewPanel.webview.html = pageHTML
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


const svgb64 = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGNsYXNzPSJwaWtjaHIiIHZpZXdCb3g9IjAgMCA0MjMuODIxIDIxNy40NCI+Cjxwb2x5Z29uIHBvaW50cz0iMTQ2LDM3IDEzNCw0MSAxMzQsMzMiIHN0eWxlPSJmaWxsOnJnYigwLDAsMCkiPjwvcG9seWdvbj4KPHBhdGggZD0iTTIsMzdMMTQwLDM3IiBzdHlsZT0iZmlsbDpub25lO3N0cm9rZS13aWR0aDoyLjE2O3N0cm9rZTpyZ2IoMCwwLDApOyI+PC9wYXRoPgo8dGV4dCB4PSI3NCIgeT0iMjUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZpbGw9InJnYigwLDAsMCkiIGRvbWluYW50LWJhc2VsaW5lPSJjZW50cmFsIj5NYXJrZG93bjwvdGV4dD4KPHRleHQgeD0iNzQiIHk9IjQ5IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJyZ2IoMCwwLDApIiBkb21pbmFudC1iYXNlbGluZT0iY2VudHJhbCI+U291cmNlPC90ZXh0Pgo8cGF0aCBkPSJNMTYxLDcyTDI1OCw3MkExNSAxNSAwIDAgMCAyNzMgNTdMMjczLDE3QTE1IDE1IDAgMCAwIDI1OCAyTDE2MSwyQTE1IDE1IDAgMCAwIDE0NiAxN0wxNDYsNTdBMTUgMTUgMCAwIDAgMTYxIDcyWiIgc3R5bGU9ImZpbGw6bm9uZTtzdHJva2Utd2lkdGg6Mi4xNjtzdHJva2U6cmdiKDAsMCwwKTsiPjwvcGF0aD4KPHRleHQgeD0iMjA5IiB5PSIxNyIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZmlsbD0icmdiKDAsMCwwKSIgZG9taW5hbnQtYmFzZWxpbmU9ImNlbnRyYWwiPk1hcmtkb3duPC90ZXh0Pgo8dGV4dCB4PSIyMDkiIHk9IjM3IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJyZ2IoMCwwLDApIiBkb21pbmFudC1iYXNlbGluZT0iY2VudHJhbCI+Rm9ybWF0dGVyPC90ZXh0Pgo8dGV4dCB4PSIyMDkiIHk9IjU3IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJyZ2IoMCwwLDApIiBkb21pbmFudC1iYXNlbGluZT0iY2VudHJhbCI+KG1hcmtkb3duLmMpPC90ZXh0Pgo8cG9seWdvbiBwb2ludHM9IjQxNywzNyA0MDUsNDEgNDA1LDMzIiBzdHlsZT0iZmlsbDpyZ2IoMCwwLDApIj48L3BvbHlnb24+CjxwYXRoIGQ9Ik0yNzMsMzdMNDExLDM3IiBzdHlsZT0iZmlsbDpub25lO3N0cm9rZS13aWR0aDoyLjE2O3N0cm9rZTpyZ2IoMCwwLDApOyI+PC9wYXRoPgo8dGV4dCB4PSIzNDUiIHk9IjI1IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJyZ2IoMCwwLDApIiBkb21pbmFudC1iYXNlbGluZT0iY2VudHJhbCI+SFRNTCtTVkc8L3RleHQ+Cjx0ZXh0IHg9IjM0NSIgeT0iNDkiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZpbGw9InJnYigwLDAsMCkiIGRvbWluYW50LWJhc2VsaW5lPSJjZW50cmFsIj5PdXRwdXQ8L3RleHQ+Cjxwb2x5Z29uIHBvaW50cz0iMjA5LDcyIDIxNCw4NCAyMDUsODQiIHN0eWxlPSJmaWxsOnJnYigwLDAsMCkiPjwvcG9seWdvbj4KPHBvbHlnb24gcG9pbnRzPSIyMDksMTQ0IDIwNSwxMzMgMjE0LDEzMyIgc3R5bGU9ImZpbGw6cmdiKDAsMCwwKSI+PC9wb2x5Z29uPgo8cGF0aCBkPSJNMjA5LDc4TDIwOSwxMzgiIHN0eWxlPSJmaWxsOm5vbmU7c3Ryb2tlLXdpZHRoOjIuMTY7c3Ryb2tlOnJnYigwLDAsMCk7Ij48L3BhdGg+CjxwYXRoIGQ9Ik0xNzYsMjE1TDI0MywyMTVBMTUgMTUgMCAwIDAgMjU4IDIwMEwyNTgsMTU5QTE1IDE1IDAgMCAwIDI0MyAxNDRMMTc2LDE0NEExNSAxNSAwIDAgMCAxNjEgMTU5TDE2MSwyMDBBMTUgMTUgMCAwIDAgMTc2IDIxNVoiIHN0eWxlPSJmaWxsOm5vbmU7c3Ryb2tlLXdpZHRoOjIuMTY7c3Ryb2tlOnJnYigwLDAsMCk7Ij48L3BhdGg+Cjx0ZXh0IHg9IjIwOSIgeT0iMTU5IiB0ZXh0LWFuY2hvcj0ibWlkZGxlIiBmaWxsPSJyZ2IoMCwwLDApIiBkb21pbmFudC1iYXNlbGluZT0iY2VudHJhbCI+UGlrY2hyPC90ZXh0Pgo8dGV4dCB4PSIyMDkiIHk9IjE4MCIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZmlsbD0icmdiKDAsMCwwKSIgZG9taW5hbnQtYmFzZWxpbmU9ImNlbnRyYWwiPkZvcm1hdHRlcjwvdGV4dD4KPHRleHQgeD0iMjA5IiB5PSIyMDAiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZpbGw9InJnYigwLDAsMCkiIGRvbWluYW50LWJhc2VsaW5lPSJjZW50cmFsIj4ocGlrY2hyLmMpPC90ZXh0Pgo8L3N2Zz4="
