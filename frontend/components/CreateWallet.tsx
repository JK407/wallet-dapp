'use client'

import { useState, useEffect } from 'react'
import { createWallet } from './WalletService'
import WalletMnemonic from './WalletMnemonic'
import { EyeIcon, EyeOffIcon, CheckCircle, XCircle, LockIcon } from 'lucide-react'

interface CreateWalletProps {
    onWalletCreated: () => void;
}

interface PasswordRequirement {
    label: string;
    test: (password: string) => boolean;
}

export default function CreateWallet({ onWalletCreated }: CreateWalletProps) {
    const [password, setPassword] = useState('')
    const [confirmPassword, setConfirmPassword] = useState('')
    const [error, setError] = useState('')
    const [passwordStrength, setPasswordStrength] = useState({ score: 0, label: '' })
    const [showPassword, setShowPassword] = useState(false)
    const [showConfirmPassword, setShowConfirmPassword] = useState(false)
    const [passwordBackgrounds, setPasswordBackgrounds] = useState<string[]>([])
    const [mnemonic, setMnemonic] = useState<string | null>(null)

    const passwordRequirements: PasswordRequirement[] = [
        { label: '至少8个字符', test: (pass) => pass.length >= 8 },
        { label: '包含大写字母', test: (pass) => /[A-Z]/.test(pass) },
        { label: '包含小写字母', test: (pass) => /[a-z]/.test(pass) },
        { label: '包含数字', test: (pass) => /[0-9]/.test(pass) },
        { label: '包含特殊字符', test: (pass) => /[!@#$%^&*(),.?":{}|<>]/.test(pass) },
        { label: '不含连续重复字符', test: (pass) => !/(.)\1/.test(pass) },
    ]

    const calculatePasswordStrength = (pass: string) => {
        const score = passwordRequirements.filter(req => req.test(pass)).length
        let label = ''
        if (score <= 2) label = '弱'
        else if (score <= 4) label = '中'
        else label = '强'
        return { score, label }
    }

    useEffect(() => {
        setPasswordStrength(calculatePasswordStrength(password))
    }, [password])

    useEffect(() => {
        const maxLength = Math.max(password.length, confirmPassword.length)
        const newBackgrounds = Array(maxLength).fill('').map((_, index) => {
            if (index >= password.length || index >= confirmPassword.length) {
                return 'bg-gray-200'
            }
            return password[index] === confirmPassword[index] ? 'bg-green-300' : 'bg-red-300'
        })
        setPasswordBackgrounds(newBackgrounds)
    }, [password, confirmPassword])

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        if (passwordStrength.score < passwordRequirements.length) {
            setError('密码不符合所有要求')
            return
        }
        if (password !== confirmPassword) {
            setError('密码不匹配')
            return
        }
        try {
            const response = await createWallet(password)
            if (response && response.mnemonic) {
                setMnemonic(response.mnemonic)
            } else {
                throw new Error('No mnemonic received')
            }
        } catch (err) {
            setError('创建钱包失败: ' + (err instanceof Error ? err.message : String(err)))
        }
    }

    const handleMnemonicConfirm = () => {
        onWalletCreated()
    }

    const renderPasswordInput = (id: string, value: string, onChange: (e: React.ChangeEvent<HTMLInputElement>) => void, show: boolean, setShow: (show: boolean) => void, disableCopy: boolean = false, disablePaste: boolean = false) => (
        <div className="relative">
            <label htmlFor={id} className="block text-sm font-medium text-gray-700 mb-1">
                {id === 'password' ? '密码' : '确认密码'}
            </label>
            <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                    <LockIcon className="h-5 w-5 text-gray-400" />
                </div>
                <input
                    type={show ? "text" : "password"}
                    id={id}
                    value={value}
                    onChange={onChange}
                    required
                    className="block w-full pl-10 pr-10 py-2 border-2 border-gray-300 rounded-lg focus:ring-blue-500 focus:border-blue-500 transition duration-150 ease-in-out sm:text-sm"
                    onCopy={disableCopy ? (e) => e.preventDefault() : undefined}
                    onCut={disableCopy ? (e) => e.preventDefault() : undefined}
                    onPaste={disablePaste ? (e) => e.preventDefault() : undefined}
                />
                <button
                    type="button"
                    onClick={() => setShow(!show)}
                    className="absolute inset-y-0 right-0 pr-3 flex items-center focus:outline-none"
                >
                    {show ? <EyeOffIcon className="h-5 w-5 text-gray-400" /> : <EyeIcon className="h-5 w-5 text-gray-400" />}
                </button>
            </div>
            <div className="absolute bottom-0 left-0 right-0 h-1 flex overflow-hidden">
                {passwordBackgrounds.map((bg, index) => (
                    <div
                        key={index}
                        className={`flex-1 ${bg} transition-all duration-300 ease-in-out`}
                        style={{
                            transform: `scaleX(${index < value.length ? 1 : 0})`,
                            transformOrigin: 'left'
                        }}
                    />
                ))}
            </div>
        </div>
    )

    const renderPasswordStrength = () => (
        <div className="mt-4">
            <div className="flex justify-between mb-1">
                <span className="text-sm font-medium text-gray-700">密码强度</span>
                <span className={`text-sm font-medium ${
                    passwordStrength.label === '强' ? 'text-green-500' :
                        passwordStrength.label === '中' ? 'text-yellow-500' : 'text-red-500'
                }`}>
                    {passwordStrength.label}
                </span>
            </div>
            <div className="w-full bg-gray-200 rounded-full h-2.5 overflow-hidden">
                <div
                    className={`h-2.5 rounded-full transition-all duration-500 ease-in-out ${
                        passwordStrength.label === '强' ? 'bg-green-500' :
                            passwordStrength.label === '中' ? 'bg-yellow-500' : 'bg-red-500'
                    }`}
                    style={{ width: `${(passwordStrength.score / passwordRequirements.length) * 100}%` }}
                ></div>
            </div>
        </div>
    )

    const renderPasswordRequirements = () => (
        <div className="mt-4">
            <h3 className="text-sm font-medium text-gray-700 mb-2">密码要求：</h3>
            <ul className="space-y-1 bg-gray-50 p-3 rounded-lg">
                {passwordRequirements.map((req, index) => (
                    <li
                        key={index}
                        className="flex items-center text-sm transition-opacity duration-300 ease-in-out"
                        style={{ opacity: password ? 1 : 0, transform: `translateX(${password ? 0 : -10}px)` }}
                    >
                        {req.test(password) ? (
                            <CheckCircle className="h-4 w-4 text-green-500 mr-2 flex-shrink-0" />
                        ) : (
                            <XCircle className="h-4 w-4 text-red-500 mr-2 flex-shrink-0" />
                        )}
                        <span className={`${req.test(password) ? 'text-green-700' : 'text-red-700'} break-words`}>
                            {req.label}
                        </span>
                    </li>
                ))}
            </ul>
        </div>
    )

    if (mnemonic) {
        return <WalletMnemonic mnemonic={mnemonic} onConfirm={handleMnemonicConfirm} />
    }

    return (
        <form onSubmit={handleSubmit} className="space-y-6">
            <h2 className="text-2xl font-bold text-center text-gray-800 mb-6">创建钱包</h2>
            {renderPasswordInput('password', password, (e) => setPassword(e.target.value), showPassword, setShowPassword, true, false)}
            {renderPasswordStrength()}
            {renderPasswordRequirements()}
            {renderPasswordInput('confirmPassword', confirmPassword, (e) => setConfirmPassword(e.target.value), showConfirmPassword, setShowConfirmPassword, false, true)}
            {error && (
                <p className="text-red-500 text-sm mt-2 animate-fadeIn">
                    {error}
                </p>
            )}
            <button
                type="submit"
                className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors duration-200 transform hover:scale-105 active:scale-95"
            >
                创建钱包
            </button>
        </form>
    )
}