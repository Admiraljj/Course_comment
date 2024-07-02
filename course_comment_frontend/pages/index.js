import Head from 'next/head'
export default function Home() {
    return (
        <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100">
            <Head>
                <title>个人介绍</title>
                <meta name="description" content="这是一个个人介绍页面" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main className="flex flex-col items-center justify-center w-full flex-1 px-20 text-center">
                <h1 className="text-5xl font-bold">
                    你好，我是 <span className="text-blue-600">你的名字</span>
                </h1>

                <p className="mt-3 text-2xl">
                    欢迎来到我的个人介绍页
                </p>

                <div className="flex flex-wrap items-center justify-around max-w-4xl mt-6 sm:w-full">
                    <div className="p-6 mt-6 text-left border w-96 rounded-xl shadow-md">
                        <h3 className="text-2xl font-bold">关于我 &rarr;</h3>
                        <p className="mt-4 text-xl">
                            这是一个简短的自我介绍。你可以在这里写一些关于自己的背景、兴趣爱好等信息。
                        </p>
                    </div>

                    <div className="p-6 mt-6 text-left border w-96 rounded-xl shadow-md">
                        <h3 className="text-2xl font-bold">联系方式 &rarr;</h3>
                        <p className="mt-4 text-xl">
                            电子邮件: example@example.com
                            <br />
                            电话: 123-456-7890
                        </p>
                    </div>
                </div>
            </main>

            <footer className="flex items-center justify-center w-full h-24 border-t">
                <p className="text-xl">© 2024 你的名字. 保留所有权利。</p>
            </footer>
        </div>
    )
}
