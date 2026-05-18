let currentSection = 'mothers';

function showSection(sectionId) {
    document.querySelectorAll('main > section').forEach(s => s.style.display = 'none');
    document.getElementById(sectionId).style.display = 'block';
    currentSection = sectionId;
    loadData();
}

async function loadData() {
    const listEl = document.getElementById(`${currentSection}-list`);
    if (!listEl) return;
    
    const res = await fetch(`/api/${currentSection}`);
    const data = await res.json();
    
    listEl.innerHTML = data.map(item => `
        <tr>
            <td>${item.first_name || item.name}</td>
            <td>${item.last_name || ''}</td>
            <td>${item.email || '-'}</td>
            <td>${item.account_status || item.status || 'Active'}</td>
            <td>
                <button onclick="editItem('${currentSection}', ${item.mother_id || item.leader_id || item.group_id || item.application_id})">Edit</button>
            </td>
        </tr>
    `).join('');
}

function openMotherForm() {
    const modal = document.getElementById('form-modal');
    const fields = document.getElementById('form-fields');
    document.getElementById('modal-title').innerText = 'Add Mother';
    
    fields.innerHTML = `
        <input type="text" name="first_name" placeholder="First Name" required>
        <input type="text" name="last_name" placeholder="Last Name" required>
        <input type="email" name="email" placeholder="Email" required>
    `;
    
    modal.showModal();
}

function closeModal() {
    document.getElementById('form-modal').close();
}

async function handleFormSubmit(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const data = Object.fromEntries(formData.entries());
    
    await fetch(`/api/${currentSection}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    });
    
    closeModal();
    loadData();
}

window.onload = () => loadData();