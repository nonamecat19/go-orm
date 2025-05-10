// JavaScript functions for handling nullable fields
function toggleNullField(checkbox, fieldName) {
    const container = document.getElementById('input-container-' + fieldName);
    const isNullField = document.getElementById(fieldName + '-is-null');
    
    if (checkbox.checked) {
        container.classList.add('opacity-50', 'pointer-events-none');
        isNullField.value = "true";
        
        // Find all inputs in the container and disable them
        const inputs = container.querySelectorAll('input, select, textarea');
        inputs.forEach(input => {
            input.disabled = true;
        });
    } else {
        container.classList.remove('opacity-50', 'pointer-events-none');
        isNullField.value = "false";
        
        // Re-enable all inputs
        const inputs = container.querySelectorAll('input, select, textarea');
        inputs.forEach(input => {
            input.disabled = false;
        });
    }
}