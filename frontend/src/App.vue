<script setup>
import {reactive, onMounted} from 'vue'
// import window from '../wailsjs/go/backend/App'

const data = reactive({
	version: "0.0.1",
	demoDir: "",
	useExternel: false,
	autoDownload: false,
	shareCode: 'steam://rungame/730/76561202255233023/+csgo_download_match%20CSGO-OXk9B-Qucbf-Q8MYy-jAMWD-My6SA',
	parsing: false,
})

const getCFG = async () => {
	const cfg = await window.go.backend.App.GetCFG()
	data.version = cfg.version
	data.demoDir = cfg.demoDir
	data.useExternel = cfg.useExternel
	data.autoDownload = cfg.autoDownload
}

const setCFG = async () => {
	let cfg = {
		version: data.version,
		demoDir: data.demoDir,
		useExternel: data.useExternel,
		autoDownload: data.autoDownload,
	}
	await window.go.backend.App.SetCFG(cfg)
}

const windowMinimise = async () => {
	await window.runtime.WindowMinimise()
}

const quit = async () => {
	window.runtime.Quit()
}

const parseShareCode = async () => {
	if (data.parsing === true) return;
	data.parsing = true
	await window.go.backend.App.ParseShareCode(data.shareCode)
	data.parsing = false
}

onMounted( async () => {
	window.runtime.EventsOn('shutdown', async () => {
		console.log("收到shutdown")
		await setCFG()
	})
	
	window.runtime.EventsOn('getCFG', async () => {
		console.log("收到getCFG")
		await getCFG()
	})
	
	window.runtime.EventsOn('inited', async () => {
		console.log("收到inited")
		await getCFG()
	})
	
	await window.go.backend.App.FrontendInited()
	
	// setInterval(()=>{
	//   setCFG()
	// }, 1000)
})

</script>

<template  >
	<div class="app-main">
		<!--		-->
		<div class="navbar" data-wails-drag >
			<div style="flex-grow: 2;" ></div>
			<a style="cursor: default" >CSGO Demo 下载器</a>
			<div style="flex-grow: 1.5; justify-content: right; display: flex; gap: .5rem;" >
				<button class="title-btn" @click="windowMinimise()">-</button>
				<button class="title-btn" @click="quit()">x</button>
			</div>
		</div>
		
		<!-- 链接区 -->
		<section class="sharecode" >
			<textarea class="texts" v-model="data.shareCode"  ></textarea>
		</section>
		
		<!-- 功能区 -->
		<section class="func" >
<!--			<div>{{data.useExternel}}</div>-->
<!--			<div>{{data.autoDownload}}</div>-->

			<div style="flex-grow: 1">版本号：{{data.version}}</div>
<!--			<button class="btn" disabled >自动下载</button>-->
<!--			<button class="btn" disabled >内置下载</button>-->
			<button class="btn" @click="parseShareCode" >
				<span v-if="!data.parsing" >解析分享链接</span>
				<span v-else>解析中...</span>
			</button>
		</section>
	</div>
</template>

<style>
.app-main {
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	//gap: 1rem;
	
	width: 100vw;
	height: 100vh;
	overflow: hidden;
}

.navbar {
	height: 3rem;
	width: 98%;
	
	display: flex;
	flex-direction: row;
	justify-content: center;
	align-items: center;
}

.sharecode {
	flex-grow: 1;
	width: 100vw;
	//padding: 0 1rem;
	display: flex;
	//border: 1px red solid;
}

.texts {
	flex-grow: 1;
	padding: 1rem;
	width: 100vw;
	//height: 100%;
	outline: none;
	border: none;
	font-size: large;
	color: white;
	background: transparent;
	resize: none;
	//font-weight: bold;
}

.func {
	width: 95%;
	height: 5rem;
	display: flex;
	flex-direction: row;
	justify-content: center;
	align-items: center;
	gap: 1rem;
}

.btn {
	padding: .75rem 1rem;
	border-radius: .25rem;
	border: none;
	outline: none;
	background: rgba(255,255,255, .3);
	font-size: large;
	color: rgba(255,255,255, .8);
	transition: all 300ms ease;
	cursor: pointer;
}

.btn:hover {
	background: rgba(255,255,255, .5);
}

.title-btn {
	width: 2rem;
	height: 2rem;
	border-radius: .35rem;
	border: none;
	outline: none;
	background: rgba(255,255,255, .3);
	font-size: large;
	color: rgba(255,255,255, .8);
	transition: all 300ms ease;
	cursor: pointer;
}

.title-btn:hover {
	background: rgba(255,255,255, .5);
}
</style>
