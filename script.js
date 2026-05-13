document.addEventListener('DOMContentLoaded', () => {
    const messageElement = document.getElementById('message');
    const messageInput = document.getElementById('message-input');
    const fontSelect = document.getElementById('font-select');
    const colorPicker = document.getElementById('color-picker');
    const messageError = document.getElementById('message-error');

    // Load saved settings or apply defaults
    messageElement.textContent = localStorage.getItem('message') || messageInput.value;
    messageElement.style.fontFamily = localStorage.getItem('fontFamily') || fontSelect.value;
    messageElement.style.color = localStorage.getItem('color') || colorPicker.value;

    messageInput.addEventListener('input', () => {
        if (messageInput.value.trim() === '') {
            messageError.textContent = 'Message cannot be empty!';
            messageElement.textContent = 'Wow Mom!'; // Revert to default or previous valid message
        } else {
            messageError.textContent = '';
            messageElement.textContent = messageInput.value;
            localStorage.setItem('message', messageInput.value);
        }
    });

    fontSelect.addEventListener('change', () => {
        messageElement.style.fontFamily = fontSelect.value;
        localStorage.setItem('fontFamily', fontSelect.value);
    });

    colorPicker.addEventListener('input', () => {
        messageElement.style.color = colorPicker.value;
        localStorage.setItem('color', colorPicker.value);
    });
});
