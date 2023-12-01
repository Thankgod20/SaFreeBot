const Web3 = require('web3');
const ganache = require("ganache");
const request = require('request');
const prompt = require('prompt');
const PancakeFactory = require("../build/contract/pancakeFactory.json");
const PancakeRouter = require("../build/contract/pancakeRouter.json");
const ERC20 = require("../build/contract/ERC20.json");
const axios = require('axios');
const qs = require('qs');
const fs = require('fs');
const process = require('process');
const HDWalletProvider = require('@truffle/hdwallet-provider');
//const mnemonic = require("../secrets.json").mnemonic;

class Safree {
    constructor() {
        this.provider = null;
        this.web3 = null;
        this.options =  null;
        this.report = {"Report":[]};
        this.error =""
        this.pair = ""
        this.whale=""
        this.nxtAct = 3
        this.address = null;
        this.maxTokenBalance = null;
        this.chatId = null;
        this.isChecked = 0;
        this.isAddress= 0;
        this.isWhale= 0;
        this.dummyS = 0;
        this.addressZero = null;
    }
    main(DexRouter,path,num,creator,rpc) {
        const initiate = async(DexRouter,path,num,creator,rpc)=>{
            return new Promise(async(resolve,reject)=>{
                this.options = {"default_balance_ether":2000,"fork":rpc,"wallet":{"unlockedAccounts":[creator,"0x0000000000000000000000000000000000000000","0x000000000000000000000000000000000000dead"]}};
                this.provider = ganache.provider(this.options);
                this.web3 = new Web3(this.provider);
                let address  = await this.web3.eth.getAccounts().catch(e=>reject(e))
                this.address = address

                let report = await swap(this.web3,this.address,path,DexRouter,num).catch(err=>{reject("Error: Equv-Amount");/*console.log("E-RR",err)*/});
                //resolve(report)

                let SwapBack = await swapBack(this.web3,this.address,path,DexRouter,num,report[0]).catch(err=>{reject("Error: Equv-Amount");/*console.log("E-RR",err)*/});

                resolve([report[1],SwapBack[1]])

            })
        }
        const swap = async(web3,address,path,DexRouter,num) => {
            return new Promise(async(resolve,reject)=>{
                //console.log("Path",path)
                let amountIn = 0.01
                let val = web3.utils.toWei(amountIn.toString(),'ether')
                let router = new web3.eth.Contract(
                    PancakeRouter.abi,
                    DexRouter
                );
                let BUSD_ERC20 = new web3.eth.Contract(
                    ERC20.abi,
                    path[(path.length)-1]
                );
                
                let swap = await router.methods.swapExactETHForTokensSupportingFeeOnTransferTokens(
                    0,
                    path,
                    address[num],
                    Math.floor(Date.now() / 1000) + 60 * 10).send({from:address[num],value:val,gas:600000,gasPrice:10000000000}).catch(e=>{/*console.log(e);*/return reject(e)});


                //Calculate Tax
                let amountTokenIn0 = await BUSD_ERC20.methods.balanceOf(address[num]).call({from:address[num]}).catch(err=>{/*console.log("E-RR",err);*/lereject("Error: Transfer_From_Failed");});

                let equv = await router.methods.getAmountsOut(amountTokenIn0,path.reverse()).call({from:address[num]}).catch(err=>{reject("Error: Equv-Amount");/*console.log("E-RR",err)*/});
                let tax = ((parseInt(val)-parseInt(equv[(equv.length)-1]))/parseInt(val))*100
                //console.log(equv,val,tax.toFixed(2))
                resolve([equv[(equv.length)-1],tax.toFixed(2)])

            })
        }
        const swapBack = async(web3,address,path,DexRouter,num,EthEquv) => {
            return new Promise(async(resolve,reject)=>{
                //console.log("Path",path)
                let router = new web3.eth.Contract(
                    PancakeRouter.abi,
                    DexRouter
                );
                let BUSD_ERC20 = new web3.eth.Contract(
                    ERC20.abi,
                    path[0]
                );
                let amountTokenIn0 = await BUSD_ERC20.methods.balanceOf(address[num]).call({from:address[num]}).catch(err=>{/*console.log("E-RR",err);*/reject("Error:Get Balance Faied");});


                let approve = await BUSD_ERC20.methods.approve(DexRouter,amountTokenIn0).send({from:address[num]}).catch(err=>{/*console.log("E-RR",err);*/reject("Error: Approved Faied");});

                

                let swap = await router.methods.swapExactTokensForETHSupportingFeeOnTransferTokens(
                                            amountTokenIn0,
                                            0,
                                            path,
                                            address[num],
                                            Math.floor(Date.now() / 1000) + 60 * 10).send({from:address[num],gas:600000,gasPrice:10000000000}).catch(err=>{/*console.log("E-RR",err);*/reject("Error:SwapBack Failed");});


                
                

                let equv = await router.methods.getAmountsOut(EthEquv,path.reverse()).call({from:address[num]}).catch(err=>{reject("Error: Equv-Amount");//console.log("E-RR",err)
            });
                let tax = ((parseInt(amountTokenIn0)-parseInt(equv[(equv.length)-1]))/parseInt(amountTokenIn0))*100
                //console.log(equv,amountTokenIn0,tax.toFixed(2))
    
                resolve([equv[(equv.length)-1],tax.toFixed(2)])
            })
        }

        return new Promise(async(resolve,reject)=>{
            initiate(DexRouter,path,num,creator,rpc).then(result=>{
                console.log("Taxe",result)
            })
        })
    }
}

if (process.argv.length<7) {
    
    console.log("node swapTest Dex [path] num owner rpc")
} else {
    
    let path = process.argv[3].split(",")
    //console.log("Path",path)
    new Safree().main(process.argv[2],path,process.argv[4],process.argv[5],process.argv[6]).then(result =>{ console.log("Result",result);}); 
}
//0x10ED43C718714eb63d5aA57B78B54704E256024E 0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c,0x1dcDC54cFd22E0FF5586A505c827D55A6D8ceB1d 2 0xB0A478255452F7D7401dE860415cC1038113a8eA https://bsc-dataseed1.binance.org/