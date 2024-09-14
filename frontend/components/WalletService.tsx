// const API_URL = 'http://127.0.0.1:8888'
const PRE_FIX = 'api/v1/wallet'

export function checkWalletExists(): boolean {
    return localStorage.getItem('wallet') !== null
}

export async function createWallet(password: string): Promise<{ mnemonic: string }> {
    const response = await fetch(`http://127.0.0.1:8888/${PRE_FIX}/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ password }),
    })

    if (!response.ok) {
        throw new Error('Failed to create wallet')
    }

    const resp = await response.json(); // 解析JSON数据
    // 现在根据data中的code属性判断是否创建成功
    if (resp.code === 2000) {
        // 成功逻辑
        localStorage.setItem('wallet', JSON.stringify({
            id: resp.data.wallet_id,
            address: resp.data.address // 假设response中也包含了address信息
        }));
        console.log('创建钱包成功, 本地存储信息:', localStorage.getItem('wallet'));
        console.log("resp:", resp)
        return { mnemonic: resp.data.mnemonic };
    } else {
        // 失败逻辑
        console.error('创建钱包失败, 服务器返回的错误信息:', resp.message);
        throw new Error('Failed to create wallet: ' + resp.message); // 抛出具体的错误信息
    }
}

export async function loginWallet(password: string): Promise<void> {
    const wallet  = localStorage.getItem('wallet');
    console.log("wallet from localStorage:", wallet);

    if (!wallet) {
        throw new Error('No wallet found');
    }

    // 将 wallet_id 转换为数字，并进行类型断言
    const { id, address } = JSON.parse(wallet);
    if (!id || !address) {
        throw new Error('Invalid wallet data');
    }

    console.log("Using wallet ID for login:", id);

    const response = await fetch(`http://127.0.0.1:8888/${PRE_FIX}/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ wallet_id:id, password, address }),
    });

    if (!response.ok) {
        throw new Error('Failed to login');
    }

    // 可以在这里处理登录成功后的逻辑,例如保存 token 等
    return await response.json();
}

