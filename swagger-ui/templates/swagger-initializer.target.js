window.onload = function() {
  window.ui = SwaggerUIBundle({
    dom_id: '#swagger-ui',
    urls: blankToUndefinedObject('{{ .URLs }}'),
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.TopBar,
    ],
    layout: "StandaloneLayout"
  });
  //</editor-fold>
};


function blankToUndefined(input) {
  return (input || '').trim() === '' ? undefined : input
}

function blankToUndefinedNumber(input) {
  let cleanInput = blankToUndefined(input)
  if (!cleanInput) {
    return undefined
  }

  const parsed = parseInt(cleanInput, 10);
  if (isNaN(parsed)) {
    return undefined;
  }

  return parsed
}

function blankToUndefinedBool(input) {
  if (input === 'true') return true;
  if (input === 'false') return false;
  return undefined
}

function blankToUndefinedArray(input) {
  const arr = input.split(",").map(blankToUndefined).filter(e => !!e)
  if ((arr || []).length === 0) {
    return undefined
  }

  return arr
}

function blankToUndefinedObject(input) {
  if (!input) {
    return undefined
  }

  input = input.trim()
  if (input.length === 0) {
    return undefined
  }

  input = decodeBase64(input);

  return JSON.parse(input)
}

function decodeBase64(str) {
  if (!str) {
    return undefined
  }

  const percentEncodedStr = atob(str || '').split('').map(c => {
    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
  }).join('');

  return decodeURIComponent(percentEncodedStr);
}


function parseJson(str) {
  if (!str) {
    return undefined
  }

  return JSON.parse(str);
}

