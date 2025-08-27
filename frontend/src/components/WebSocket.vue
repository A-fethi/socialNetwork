<template>
    <div>
        <h1>WebSocket</h1>
        <div>
            <div>Receiver: {{ receiver }}</div>
        </div>
        <div>
            <input v-model="message" @keyup.enter="sendMessage" placeholder="Type a message..." />
            <button @click="sendMessage">Send</button>
        </div>
        <div v-for="(msg, index) in messages" :key="index">{{ msg }}</div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            socket: null,
            message: '',
            receiver: 'admin',
            messages: [],
            currentMessage: {
                type: '',
                message: '',
                receiver: '',
                username: ''
            }
        };
    },
    
    mounted() {
        
        this.socket = new WebSocket("ws://localhost:8080/ws");

        this.socket.onmessage = (event) => {
            const data = JSON.parse(event.data);
           console.log(data);
           
        };

        this.socket.onopen = () => {
            console.log("WebSocket connected");
        };

        this.socket.onerror = (error) => {
            console.error("WebSocket error:", error);
        };

        this.socket.onclose = () => {
            console.log("WebSocket disconnected");
        };
    },
    methods: {
        sendMessage() {
            if (this.socket && this.message.trim() !== "") {
                this.currentMessage.type = "message";
                this.currentMessage.receiver = this.receiver;
                this.currentMessage.username = "afethi";
                this.currentMessage.message = this.message;
                this.socket.send(JSON.stringify(this.currentMessage));
                this.messages.push(`Sent: ${this.message}`);
                this.message = "";
            }
        },
    },
};
</script>