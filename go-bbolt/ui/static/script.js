/* static/script.js */
document.addEventListener('DOMContentLoaded', () => {
    hljs.highlightAll();
});

async function toggleValue(element, bucket, key) {
    const valueDiv = element.nextElementSibling;
    const toggleSpan = element.querySelector('.toggle');
    
    if (valueDiv.classList.contains('hidden')) {
        if (!valueDiv.dataset.loaded) {
            const response = await fetch(`/api/bucket/${bucket}/key/${key}`);
            const data = await response.json();
            
            if (data.isJSON) {
                valueDiv.innerHTML = `<pre><code class="language-json">${JSON.stringify(data.jsonValue, null, 2)}</code></pre>`;
                hljs.highlightElement(valueDiv.querySelector('code'));
            } else {
                valueDiv.textContent = data.value;
            }
            valueDiv.dataset.loaded = 'true';
        }
        valueDiv.classList.remove('hidden');
        toggleSpan.classList.add('open');
    } else {
        valueDiv.classList.add('hidden');
        toggleSpan.classList.remove('open');
    }
}