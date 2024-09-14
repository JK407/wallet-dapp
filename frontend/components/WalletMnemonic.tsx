'use client'

import { useState } from 'react'
import { EyeIcon, EyeOffIcon, ClipboardIcon } from 'lucide-react'
import { motion } from 'framer-motion'

interface WalletMnemonicProps {
    mnemonic: string;
    onConfirm: () => void;
}

export default function WalletMnemonic({ mnemonic, onConfirm }: WalletMnemonicProps) {
    const [showMnemonic, setShowMnemonic] = useState(false)
    const [copied, setCopied] = useState(false)

    const handleCopy = () => {
        navigator.clipboard.writeText(mnemonic)
        setCopied(true)
        setTimeout(() => setCopied(false), 2000)
    }

    return (
        <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5 }}
            className="space-y-6"
        >
            <h2 className="text-2xl font-bold text-center text-gray-800 mb-6">您的钱包助记词</h2>
            <p className="text-sm text-gray-600 mb-4">
                请将以下助记词安全地保存下来。您将需要它来恢复您的钱包。切勿与他人分享！
            </p>
            <div className="relative bg-gray-100 p-4 rounded-lg mb-6">
                <motion.div
                    initial={{ filter: 'blur(4px)' }}
                    animate={{ filter: showMnemonic ? 'blur(0px)' : 'blur(4px)' }}
                    transition={{ duration: 0.3 }}
                    className="break-words text-gray-800"
                >
                    {mnemonic}
                </motion.div>
                {!showMnemonic && (
                    <div className="absolute inset-0 flex items-center justify-center bg-gray-200 bg-opacity-50 backdrop-blur-sm rounded-lg">
                        点击下方按钮显示助记词
                    </div>
                )}
            </div>
            <div className="flex justify-between mb-6">
                <motion.button
                    whileHover={{ scale: 1.05 }}
                    whileTap={{ scale: 0.95 }}
                    type="button"
                    onClick={() => setShowMnemonic(!showMnemonic)}
                    className="flex items-center px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                >
                    {showMnemonic ? (
                        <>
                            <EyeOffIcon className="h-5 w-5 mr-2" />
                            隐藏助记词
                        </>
                    ) : (
                        <>
                            <EyeIcon className="h-5 w-5 mr-2" />
                            显示助记词
                        </>
                    )}
                </motion.button>
                <motion.button
                    whileHover={{ scale: 1.05 }}
                    whileTap={{ scale: 0.95 }}
                    type="button"
                    onClick={handleCopy}
                    className="flex items-center px-4 py-2 border border-gray-300 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                >
                    <ClipboardIcon className="h-5 w-5 mr-2" />
                    {copied ? '已复制' : '复制'}
                </motion.button>
            </div>
            <motion.button
                whileHover={{ scale: 1.05 }}
                whileTap={{ scale: 0.95 }}
                onClick={onConfirm}
                className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors duration-200"
            >
                我已安全保存助记词
            </motion.button>
        </motion.div>
    )
}