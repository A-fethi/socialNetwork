<template>
<div  id="info">
    <h1>Loading Information</h1>
</div>

</template>
<script>

export default{
    name: 'ChechInfo',
    data() {
        return {
            info: null,
        };
    },
    methods: {
        async fetchInfo() {
            try {
                const response = await fetch('http://localhost:8080/api/info', {
                    method: 'GET',
                    credentials: 'include',
                });

                if (!response.ok) {
                    const errorText = await response.text();
                    throw new Error(errorText || 'Network response was not ok');
                }

                const data = await response.json();
                console.log('Data:', data);
                
                this.info = data;
            } catch (error) {
                console.error('Error fetching info:', error);
                this.showNotification(error.message || 'Error fetching info', 'error');
            }
        },
    },
    created() {
    this.fetchInfo();
  },
}

</script>