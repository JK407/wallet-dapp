'use client'

import { useState } from 'react'
import { loginWallet } from './WalletService'
import { EyeIcon, EyeOffIcon, LockIcon } from 'lucide-react'
import { motion } from 'framer-motion'

export default function LoginWithPassword() {
    const [password, setPassword] = useState('')
    const [error, setError] = useState('')
    const [showPassword, setShowPassword] = useState(false)

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            const response = await loginWallet(password)
            alert(JSON.stringify(response, null, 2))
            console.log('登录成功,完整response:', response)
            // 登录成功后的处理,例如跳转到钱包主页
        } catch (err) {
            setError('登录失败,请检查密码')
        }
    }

    return (
        <motion.form
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5 }}
            onSubmit={handleSubmit}
            className="space-y-6"
        >
            <div className="relative">
                <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">
                    密码
                </label>
                <div className="relative">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <LockIcon className="h-5 w-5 text-gray-400" />
                    </div>
                    <input
                        type={showPassword ? "text" : "password"}
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                        className="block w-full pl-10 pr-10 py-2 border-2 border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500 transition duration-150 ease-in-out sm:text-sm"
                    />
                    <button
                        type="button"
                        onClick={() => setShowPassword(!showPassword)}
                        className="absolute inset-y-0 right-0 pr-3 flex items-center focus:outline-none"
                    >
                        {showPassword ? <EyeOffIcon className="h-5 w-5 text-gray-400" /> : <EyeIcon className="h-5 w-5 text-gray-400" />}
                    </button>
                </div>
            </div>
            {error && (
                <motion.p
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    className="text-red-500 text-sm mt-2"
                >
                    {error}
                </motion.p>
            )}
            <motion.button
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
                type="submit"
                className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors duration-200"
            >
                登录钱包
            </motion.button>
        </motion.form>
    )
}