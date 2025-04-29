// src/api/fetchInfo.js
export async function fetchInfo() {
    try {
      const response = await fetch('http://localhost:8080/api/info', {
        method: 'GET',
        credentials: 'include',
      });
  
      if (!response.ok) {
        throw new Error('Failed to fetch user info');
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error fetching user info:', error);
      throw error;
    }
  }
  