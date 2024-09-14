'use client'

import { useState, useEffect } from 'react'
import CreateWallet from './CreateWallet'
import LoginWithPassword from './LoginWithPassword'
import { checkWalletExists } from './WalletService'
import { motion } from 'framer-motion'

export default function LoginPage() {
  const [hasWallet, setHasWallet] = useState<boolean | null>(null)

  useEffect(() => {
    const walletExists = checkWalletExists()
    setHasWallet(walletExists)
  }, [])

  const handleWalletCreated = () => {
    setHasWallet(true)
  }

  if (hasWallet === null) {
    return (
        <div className="min-h-screen flex items-center justify-center bg-gradient-to-r from-blue-500 to-purple-600">
          <motion.div
              initial={{ opacity: 0, scale: 0.9 }}
              animate={{ opacity: 1, scale: 1 }}
              transition={{ duration: 0.5 }}
              className="bg-white p-8 rounded-lg shadow-2xl w-full max-w-md"
          >
            <div className="flex justify-center items-center space-x-2">
              <div className="w-4 h-4 bg-blue-500 rounded-full animate-bounce"></div>
              <div className="w-4 h-4 bg-blue-500 rounded-full animate-bounce delay-100"></div>
              <div className="w-4 h-4 bg-blue-500 rounded-full animate-bounce delay-200"></div>
            </div>
          </motion.div>
        </div>
    )
  }

  return (
      <div className="min-h-screen flex items-center justify-center bg-gradient-to-r from-blue-500 to-purple-600 p-4">
        <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.5 }}
            className="bg-white p-8 rounded-lg shadow-2xl w-full max-w-md"
        >
          <h1 className="text-3xl font-bold mb-6 text-center text-gray-800">区块链钱包</h1>
          {hasWallet ? (
              <LoginWithPassword />
          ) : (
              <CreateWallet onWalletCreated={handleWalletCreated} />
          )}
        </motion.div>
      </div>
  )
}