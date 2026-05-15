// Utility function to make API requests
async function apiRequest(method, url, data = null) {
    try {
        const options = {
            method,
            headers: {
                'Content-Type': 'application/json',
            },
        };
        if (data) {
            options.body = JSON.stringify(data);
        }
        const response = await fetch(url, options);
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`${response.status} ${response.statusText}: ${errorText}`);
        }
        if (method === 'DELETE') {
            return null; // No content for delete
        }
        return await response.json();
    } catch (error) {
        console.error(`API Request Failed (${method} ${url}):`, error);
        alert(`Error: ${error.message}`);
        throw error; // Re-throw to propagate the error
    }
}

// Generic function to load data into tables
async function loadTableData(tableName, apiUrl, tableBodyId, renderRowFn) {
    try {
        const data = await apiRequest('GET', apiUrl);
        const tableBody = document.getElementById(tableBodyId);
        tableBody.innerHTML = ''; // Clear existing data
        data.forEach(item => {
            const row = renderRowFn(item);
            tableBody.appendChild(row);
        });
    } catch (error) {
        console.error(`Failed to load ${tableName} data:`, error);
    }
}

// --- Mothers CRUD Operations ---
const motherModal = document.getElementById('mother-modal');
const motherForm = document.getElementById('mother-form');

function openMotherModal(mother = {}) {
    document.getElementById('mother-modal-title').textContent = mother.mother_id ? 'Edit Mother' : 'Add Mother';
    document.getElementById('mother-id').value = mother.mother_id || '';
    document.getElementById('mother-first-name').value = mother.first_name || '';
    document.getElementById('mother-last-name').value = mother.last_name || '';
    document.getElementById('mother-email').value = mother.email || '';
    document.getElementById('mother-phone-number').value = mother.phone_number || '';
    document.getElementById('mother-address').value = mother.address || '';
    document.getElementById('mother-city').value = mother.city || '';
    document.getElementById('mother-state').value = mother.state || '';
    document.getElementById('mother-zip-code').value = mother.zip_code || '';
    document.getElementById('mother-date-of-birth').value = mother.date_of_birth ? mother.date_of_birth.split('T')[0] : '';
    document.getElementById('mother-children-info').value = mother.children_info || '';
    document.getElementById('mother-preferred-meeting-times').value = mother.preferred_meeting_times || '';
    document.getElementById('mother-preferred-locations').value = mother.preferred_locations || '';
    motherModal.style.display = 'block';
}

function closeMotherModal() {
    motherModal.style.display = 'none';
    motherForm.reset();
}

motherForm.addEventListener('submit', async (event) => {
    event.preventDefault();
    const motherId = document.getElementById('mother-id').value;
    const mother = {
        first_name: document.getElementById('mother-first-name').value,
        last_name: document.getElementById('mother-last-name').value,
        email: document.getElementById('mother-email').value,
        phone_number: document.getElementById('mother-phone-number').value,
        address: document.getElementById('mother-address').value,
        city: document.getElementById('mother-city').value,
        state: document.getElementById('mother-state').value,
        zip_code: document.getElementById('mother-zip-code').value,
        date_of_birth: document.getElementById('mother-date-of-birth').value,
        children_info: document.getElementById('mother-children-info').value,
        preferred_meeting_times: document.getElementById('mother-preferred-meeting-times').value,
        preferred_locations: document.getElementById('mother-preferred-locations').value,
    };

    try {
        if (motherId) {
            await apiRequest('PUT', `/api/mothers/${motherId}`, mother);
        } else {
            await apiRequest('POST', '/api/mothers', mother);
        }
        closeMotherModal();
        loadMothers();
    } catch (error) {
        console.error('Failed to save mother:', error);
    }
});

async function deleteMother(motherId) {
    if (confirm('Are you sure you want to delete this mother?')) {
        try {
            await apiRequest('DELETE', `/api/mothers/${motherId}`);
            loadMothers();
        } catch (error) {
            console.error('Failed to delete mother:', error);
        }
    }
}

function renderMotherRow(mother) {
    const row = document.createElement('tr');
    row.innerHTML = `
        <td>${mother.mother_id}</td>
        <td>${mother.first_name}</td>
        <td>${mother.last_name}</td>
        <td>${mother.email}</td>
        <td>${mother.phone_number || ''}</td>
        <td>${mother.city || ''}</td>
        <td>${mother.state || ''}</td>
        <td>${mother.account_status || 'Active'}</td>
        <td class="action-buttons">
            <button onclick="openMotherModal(${JSON.stringify(mother).replace(/'/g, "&apos;")})">Edit</button>
            <button class="delete" onclick="deleteMother(${mother.mother_id})">Delete</button>
        </td>
    `;
    return row;
}

function loadMothers() {
    loadTableData('Mothers', '/api/mothers', 'mothers-table-body', renderMotherRow);
}

// --- Leaders CRUD Operations ---
const leaderModal = document.getElementById('leader-modal');
const leaderForm = document.getElementById('leader-form');

function openLeaderModal(leader = {}) {
    document.getElementById('leader-modal-title').textContent = leader.leader_id ? 'Edit Leader' : 'Add Leader';
    document.getElementById('leader-id').value = leader.leader_id || '';
    document.getElementById('leader-mother-id').value = leader.mother_id || '';
    document.getElementById('leader-first-name').value = leader.first_name || '';
    document.getElementById('leader-last-name').value = leader.last_name || '';
    document.getElementById('leader-email').value = leader.email || '';
    document.getElementById('leader-phone-number').value = leader.phone_number || '';
    document.getElementById('leader-bio').value = leader.bio || '';
    document.getElementById('leader-status').value = leader.leader_status || 'Pending Approval';
    document.getElementById('leader-approved-by-admin').checked = leader.approved_by_admin || false;
    leaderModal.style.display = 'block';
}

function closeLeaderModal() {
    leaderModal.style.display = 'none';
    leaderForm.reset();
}

leaderForm.addEventListener('submit', async (event) => {
    event.preventDefault();
    const leaderId = document.getElementById('leader-id').value;
    const leader = {
        mother_id: document.getElementById('leader-mother-id').value ? parseInt(document.getElementById('leader-mother-id').value) : null,
        first_name: document.getElementById('leader-first-name').value,
        last_name: document.getElementById('leader-last-name').value,
        email: document.getElementById('leader-email').value,
        phone_number: document.getElementById('leader-phone-number').value,
        bio: document.getElementById('leader-bio').value,
        leader_status: document.getElementById('leader-status').value,
        approved_by_admin: document.getElementById('leader-approved-by-admin').checked,
    };

    try {
        if (leaderId) {
            await apiRequest('PUT', `/api/leaders/${leaderId}`, leader);
        } else {
            await apiRequest('POST', '/api/leaders', leader);
        }
        closeLeaderModal();
        loadLeaders();
    } catch (error) {
        console.error('Failed to save leader:', error);
    }
});

async function deleteLeader(leaderId) {
    if (confirm('Are you sure you want to delete this leader?')) {
        try {
            await apiRequest('DELETE', `/api/leaders/${leaderId}`);
            loadLeaders();
        } catch (error) {
            console.error('Failed to delete leader:', error);
        }
    }
}

function renderLeaderRow(leader) {
    const row = document.createElement('tr');
    row.innerHTML = `
        <td>${leader.leader_id}</td>
        <td>${leader.first_name}</td>
        <td>${leader.last_name}</td>
        <td>${leader.email}</td>
        <td>${leader.phone_number || ''}</td>
        <td>${leader.leader_status || 'Pending Approval'}</td>
        <td>${leader.approved_by_admin ? 'Yes' : 'No'}</td>
        <td class="action-buttons">
            <button onclick="openLeaderModal(${JSON.stringify(leader).replace(/'/g, "&apos;")})">Edit</button>
            <button class="delete" onclick="deleteLeader(${leader.leader_id})">Delete</button>
        </td>
    `;
    return row;
}

function loadLeaders() {
    loadTableData('Leaders', '/api/leaders', 'leaders-table-body', renderLeaderRow);
}

// --- Groups CRUD Operations ---
const groupModal = document.getElementById('group-modal');
const groupForm = document.getElementById('group-form');

function openGroupModal(group = {}) {
    document.getElementById('group-modal-title').textContent = group.group_id ? 'Edit Group' : 'Add Group';
    document.getElementById('group-id').value = group.group_id || '';
    document.getElementById('group-name').value = group.name || '';
    document.getElementById('group-description').value = group.description || '';
    document.getElementById('group-leader-id').value = group.leader_id || '';
    document.getElementById('group-max-members').value = group.max_members || 15;
    document.getElementById('group-current-members').value = group.current_members || 0;
    document.getElementById('group-status').value = group.status || 'Active';
    document.getElementById('group-meeting-details').value = group.meeting_details || '';
    document.getElementById('group-meeting-frequency').value = group.meeting_frequency || '';
    groupModal.style.display = 'block';
}

function closeGroupModal() {
    groupModal.style.display = 'none';
    groupForm.reset();
}

groupForm.addEventListener('submit', async (event) => {
    event.preventDefault();
    const groupId = document.getElementById('group-id').value;
    const group = {
        name: document.getElementById('group-name').value,
        description: document.getElementById('group-description').value,
        leader_id: parseInt(document.getElementById('group-leader-id').value),
        max_members: parseInt(document.getElementById('group-max-members').value),
        current_members: parseInt(document.getElementById('group-current-members').value),
        status: document.getElementById('group-status').value,
        meeting_details: document.getElementById('group-meeting-details').value,
        meeting_frequency: document.getElementById('group-meeting-frequency').value,
    };

    try {
        if (groupId) {
            await apiRequest('PUT', `/api/groups/${groupId}`, group);
        } else {
            await apiRequest('POST', '/api/groups', group);
        }
        closeGroupModal();
        loadGroups();
    } catch (error) {
        console.error('Failed to save group:', error);
    }
});

async function deleteGroup(groupId) {
    if (confirm('Are you sure you want to delete this group?')) {
        try {
            await apiRequest('DELETE', `/api/groups/${groupId}`);
            loadGroups();
        } catch (error) {
            console.error('Failed to delete group:', error);
        }
    }
}

function renderGroupRow(group) {
    const row = document.createElement('tr');
    row.innerHTML = `
        <td>${group.group_id}</td>
        <td>${group.name}</td>
        <td>${group.description || ''}</td>
        <td>${group.leader_id}</td>
        <td>${group.max_members}</td>
        <td>${group.current_members}</td>
        <td>${group.status || 'Active'}</td>
        <td>${group.meeting_frequency || ''}</td>
        <td class="action-buttons">
            <button onclick="openGroupModal(${JSON.stringify(group).replace(/'/g, "&apos;")})">Edit</button>
            <button class="delete" onclick="deleteGroup(${group.group_id})">Delete</button>
        </td>
    `;
    return row;
}

function loadGroups() {
    loadTableData('Groups', '/api/groups', 'groups-table-body', renderGroupRow);
}

// --- Applications CRUD Operations ---
const applicationModal = document.getElementById('application-modal');
const applicationForm = document.getElementById('application-form');

function openApplicationModal(app = {}) {
    document.getElementById('application-modal-title').textContent = app.application_id ? 'Edit Application' : 'Add Application';
    document.getElementById('application-id').value = app.application_id || '';
    document.getElementById('application-mother-id').value = app.mother_id || '';
    document.getElementById('application-group-id').value = app.group_id || '';
    document.getElementById('application-status').value = app.application_status || 'Pending';
    document.getElementById('application-interview-date').value = app.interview_date ? app.interview_date.split('T')[0] : '';
    document.getElementById('application-interview-time-slot').value = app.interview_time_slot ? app.interview_time_slot.slice(0, 16) : '';
    document.getElementById('application-interview-confirmed-by-mother').checked = app.interview_confirmed_by_mother || false;
    document.getElementById('application-rejection-reason').value = app.rejection_reason || '';
    document.getElementById('application-accepted-date').value = app.accepted_date ? app.accepted_date.split('T')[0] : '';
    document.getElementById('application-activated-date').value = app.activated_date ? app.activated_date.split('T')[0] : '';
    document.getElementById('application-notes-from-leader').value = app.notes_from_leader || '';
    applicationModal.style.display = 'block';
}

function closeApplicationModal() {
    applicationModal.style.display = 'none';
    applicationForm.reset();
}

applicationForm.addEventListener('submit', async (event) => {
    event.preventDefault();
    const applicationId = document.getElementById('application-id').value;
    const app = {
        mother_id: parseInt(document.getElementById('application-mother-id').value),
        group_id: parseInt(document.getElementById('application-group-id').value),
        application_status: document.getElementById('application-status').value,
        interview_date: document.getElementById('application-interview-date').value,
        interview_time_slot: document.getElementById('application-interview-time-slot').value ? new Date(document.getElementById('application-interview-time-slot').value).toISOString() : null,
        interview_confirmed_by_mother: document.getElementById('application-interview-confirmed-by-mother').checked,
        rejection_reason: document.getElementById('application-rejection-reason').value,
        accepted_date: document.getElementById('application-accepted-date').value,
        activated_date: document.getElementById('application-activated-date').value,
        notes_from_leader: document.getElementById('application-notes-from-leader').value,
    };

    try {
        if (applicationId) {
            await apiRequest('PUT', `/api/applications/${applicationId}`, app);
        } else {
            await apiRequest('POST', '/api/applications', app);
        }
        closeApplicationModal();
        loadApplications();
    } catch (error) {
        console.error('Failed to save application:', error);
    }
});

async function deleteApplication(applicationId) {
    if (confirm('Are you sure you want to delete this application?')) {
        try {
            await apiRequest('DELETE', `/api/applications/${applicationId}`);
            loadApplications();
        } catch (error) {
            console.error('Failed to delete application:', error);
        }
    }
}

function renderApplicationRow(app) {
    const row = document.createElement('tr');
    row.innerHTML = `
        <td>${app.application_id}</td>
        <td>${app.mother_id}</td>
        <td>${app.group_id}</td>
        <td>${app.application_status || 'Pending'}</td>
        <td>${app.interview_confirmed_by_mother ? 'Yes' : 'No'}</td>
        <td class="action-buttons">
            <button onclick="openApplicationModal(${JSON.stringify(app).replace(/'/g, "&apos;")})">Edit</button>
            <button class="delete" onclick="deleteApplication(${app.application_id})">Delete</button>
        </td>
    `;
    return row;
}

function loadApplications() {
    loadTableData('Applications', '/api/applications', 'applications-table-body', renderApplicationRow);
}

// Tab switching logic
function showSection(sectionId) {
    document.querySelectorAll('.crud-section').forEach(section => {
        section.style.display = 'none';
    });
    document.getElementById(`${sectionId}-section`).style.display = 'block';
    // Load data for the active section
    switch (sectionId) {
        case 'mothers':
            loadMothers();
            break;
        case 'leaders':
            loadLeaders();
            break;
        case 'groups':
            loadGroups();
            break;
        case 'applications':
            loadApplications();
            break;
    }
}

// Initial load
window.onload = () => {
    showSection('mothers'); // Show mothers section by default
};
