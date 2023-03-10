<template>
    <div class="container">
        <div class="loginBox">
            <a-form :model="formData" :rules="rules" class="loginForm">
                <a-form-item prop="username">
                    <a-input
                        v-model:value="formData.username"
                        placeholder="Username"
                    ></a-input>
                </a-form-item>

                <a-form-item prop="password">
                    <a-input-password
                        v-model:value="formData.password"
                        placeholder="Password"
                        type="password"
                    ></a-input-password>
                </a-form-item>

                <a-form-item-model class="loginBtn">
                    <a-button type="primary" style="margin: 10px"
                        >Login</a-button
                    >
                    <a-button type="info" style="margin: 10px">Cancel</a-button>
                </a-form-item-model>
            </a-form>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            formData: {
                username: '',
                password: '',
            },
            rules: {
                username: [
                    {
                        required: true,
                        message: 'Please input username',
                        trigger: 'blur',
                    },
                    {
                        min: 4,
                        max: 12,
                        message:
                            'Username must more than 4 chars and less than 12 chars',
                        trigger: 'blur',
                    },
                ],
                password: [
                    {
                        required: true,
                        message: 'Please input password',
                        trigger: 'blur',
                    },
                    {
                        min: 6,
                        max: 20,
                        message:
                            'Password must more than 6 chars and less than 20 chars',
                        trigger: 'blur',
                    },
                ],
            },
        }
    },
}
</script>

<style scoped>
.container {
    height: 100%;
    background-color: #1fa16b;
}

.loginBox {
    width: 450px;
    height: 300px;
    background-color: #fff;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    border-radius: 8px;
}

.loginForm {
    width: 100%;
    position: absolute;
    bottom: 20%;
    padding: 0 20px;
    box-sizing: border-box;
}

.loginBtn {
    display: flex;
    justify-content: flex-end;
}
</style>
