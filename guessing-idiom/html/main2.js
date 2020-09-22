const SIZE = 8 ;
const SPAN = 30 ;
const AMOUNT = 20 ;

function Game() {
    /* 构造函数 */
}

Game.prototype = {
    /* 全局变量(不可清除) */

    // 累加关卡
    gameLevel: 0 ,

    /* 单局变量 */

    // 谜题列表
    idioms: "" ,
    idiomList: null,

    // 玩家尝试次数
    tryCounter: 0,

    // 结果集
    results: null,
    resultGrids: null,
    tmpResultRow: null,
    questObj: null,
    questEle: null,

    // 判断显示是否应该意义(兼容以后单成语多空题)
    indexIdiomToCounter: [] , // "一马当先":{c:1,q:"释义"}
    indexPosToIdiom: [] ,     // "2-3":["一马当先","马到成功"]

    // 界面元素
    eleContent: null,
    eleResult: null,
    eleAnswer: null,

    // 全局初始化
    init: function(){
        document.getElementById("next-btn").addEventListener("click", ()=>{
            this.next() ;
        }) ;

        document.getElementById("reload-btn").addEventListener("click", ()=>{
            window.location.reload() ;
        }) ;

        this.next();
    } ,

    // 清除
    cleanGame: function() {
        // 清理变量
        this.idioms = "" ;
        this.idiomList = null ;
        this.tryCounter = 0 ;
        this.results = null ;
        this.resultGrids = null ;
        this.tmpResultRow = null ;
        this.questObj = null ;
        this.questEle = null ;
        this.indexIdiomToCounter = [] ;
        this.indexPosToIdiom = [] ;

        // 清理界面
        let newContent = document.createElement("div") ;
        newContent.setAttribute("id", "content") ;
        let newBackground = document.createElement("div") ;
        newBackground.setAttribute("id", "background") ;
        newContent.append(newBackground) ;
        document.getElementById("content").replaceWith(newContent) ;

        let newResult = document.createElement("div") ;
        newResult.setAttribute("id", "result") ;
        document.getElementById("result").replaceWith(newResult) ;

        let newAnswer = document.createElement("div") ;
        newAnswer.setAttribute("id", "answer") ;
        document.getElementById("answer").replaceWith(newAnswer) ;

        // 隐藏按钮
        this.nextBtn(false) ;
    },

    // 初始化
    initGame: function() {
        // 获取元素
        this.eleAnswer = document.getElementById("answer") ;

        // 数据初始化
        this.idiomList = JSON.parse(this.idioms) ;
        this.createBackground(SIZE) ;
        let {results,resultGrids} = this.createPlayground(SPAN, this.idiomList) ;
        this.results = results ;
        this.resultGrids = resultGrids ;
        console.log("results : ", this.results) ;
        console.log("resultGrids : ", JSON.stringify(this.resultGrids)) ;
        this.renderResults(this.results) ;

        // 问题初始化
        this.questObj = this.quest() ;
        // `grid_${i}_${j}`
        this.highlightGrid(this.questObj.i, this.questObj.j) ;
    } ,

    reloadBtn: function(display) {
        let ele = document.getElementById("reload") ;
        if (display) {
            ele.setAttribute("style", "display:block") ;
        }else{
            ele.setAttribute("style", "display:none") ;
        }
    } ,

    nextBtn: function(display) {
        let ele = document.getElementById("next") ;
        if (display) {
            ele.setAttribute("style", "display:block") ;
        }else{
            ele.setAttribute("style", "display:none") ;
        }
    } ,

    showLoading: function() {
        let msg = "关卡获取中..." ;
        document.getElementById("loading").innerText = msg;
    } ,

    showFailed: function(errMsg) {
        let msg = `获取关卡失败 ${errMsg}` ;
        document.getElementById("loading").innerText = msg;
    } ,

    showDone: function() {
        let msg = `当前第 ${this.gameLevel}/${AMOUNT} 关，💪` ;
        document.getElementById("loading").innerText = msg;
    } ,

    // 下一局
    next: function() {
        this.gameLevel++ ;
        if (this.gameLevel>AMOUNT) {
            alert("恭喜通关💐");
            // window.location.reload() ;
            this.reloadBtn(true);
            return;
        }
        this.cleanGame();
        this.showLoading() ;
        fetch('/game/stage/get?level='+this.gameLevel+'&limit=1')
            .then(function(response) {
                return response.json();
            })
            .then((myJson)=> {
                if (myJson.code != 0) {
                    this.showFailed(myJson.msg);
                    return
                }
                console.log("response: ", myJson);

                if (myJson.data.list.length<=0) {
                    this.showFailed("no result returned");
                    return
                }

                this.idioms = myJson.data.list[0].Contents ;
                this.initGame() ;

                this.showDone();
            });
    } ,

    // 点击结果
    onResultClicked: function(ele) {
        this.tryCounter++ ;
        let resultIndex = `${this.questObj.i}-${this.questObj.j}` ;
        let innerText = ele.childNodes[0].nodeValue ;
        if (this.results[resultIndex]!=innerText) {
            alert('错误，请重试') ;
            return
        }

        // 更新界面
        this.questEle.setAttribute("class", "bg-normal") ;
        ele.replaceWith("");
        document.getElementById(`item_${this.questObj.i}_${this.questObj.j}`).childNodes[0].replaceWith(innerText);

        // 显示释义
        let posToIdiom = this.indexPosToIdiom[resultIndex] ;
        for (let i in posToIdiom) {
            let idiomCounter = this.indexIdiomToCounter[posToIdiom[i]] ;
            if (idiomCounter.c==0){
                continue ;
            }

            if (--idiomCounter.c == 0) {
                console.log(`【${posToIdiom[i]}】 ${idiomCounter.q}`) ;
                let ele_p = document.createElement("p") ;
                ele_p.innerText = `【${posToIdiom[i]}】 ${idiomCounter.q}` ;
                this.eleAnswer.append(ele_p) ;
            }
        }

        // 下一个迷
        this.questObj = this.quest() ;
        if (this.questObj==null) {
            if (this.gameLevel==AMOUNT) {
                alert("恭喜通关💐");
                this.reloadBtn(true);
                return;
            }else{
                alert(`回答完成：共尝试了 ${this.tryCounter} 次`) ;
                this.nextBtn(true) ;
            }
            return
        }
        this.highlightGrid(this.questObj.i, this.questObj.j) ;
    },

    // 高亮一个背景格子
    highlightGrid: function(i,j) {
        this.questEle = document.getElementById(`grid_${i}_${j}`);
        this.questEle.setAttribute("class", "bg-normal bg-highlight") ;
    },

    // 获取一个问题
    quest: function() {
        if (this.tmpResultRow==null || this.tmpResultRow.children.length<=0) {
            if (this.resultGrids.length<=0) {
                return null ;
            }
            this.tmpResultRow = this.resultGrids.shift() ;
        }

        let result = {} ;
        result.i = this.tmpResultRow.index ;
        result.j = this.tmpResultRow.children.shift() ;

        return result ;
    },

    // 结果集渲染
    renderResults: function(results) {

        let parentEle = document.getElementById("result") ;
        for (let i in results) {
            let ele = document.createElement("div") ;
            ele.setAttribute("class", "item") ;
            ele.addEventListener("click", (event)=>{
                this.onResultClicked(event.target) ;
            }) ;
            ele.append(results[i]) ;

            parentEle.appendChild(ele) ;
        }

        this.appendClean(parentEle)
    },

    // 辅助函数：添加清除样式
    appendClean: function(ele) {
        let clEle = document.createElement("div") ;
        clEle.setAttribute("class", "clean") ;
        ele.appendChild(clEle) ;
    },

    // 背景格子渲染
    createBackground: function(size) {
        let bgEle = document.getElementById("background") ;
        for (let i=0; i<size; i++) {
            let eleRow = document.createElement("div") ;
            eleRow.setAttribute("class", "bg-line") ;

            for (let j=0; j<size; j++) {
                let eleColl = document.createElement("div") ;
                eleColl.setAttribute("class", "bg-normal") ;
                eleColl.setAttribute("id", `grid_${i}_${j}`) ;

                eleRow.appendChild(eleColl) ;
            }

            bgEle.appendChild(eleRow) ;
            this.appendClean(bgEle) ;
        }
    },

    // 主游戏逻辑界面渲染
    createPlayground: function(span, data) {
        let results = [] ;      // 结果集
        let resultGrids = [] ;  // 结果集格子数组

        // 计算结果集
        for (let i in data) {
            let item = data[i] ;
            let poses = item["pos"].split(";") ;

            for (let j=0,c=item["idiom"].length; j<c; j++) {
                let pos = poses.shift() ;
                let posObj = pos.split(",") ;
                let resultsKey = `${posObj[1]}-${posObj[0]}` ;
                if (item["vis"][j]=="0") {

                    // 结果集
                    results[resultsKey] = item["idiom"][j] ;

                    // 结果格子
                    if (resultGrids[posObj[1]]==undefined) {
                        resultGrids[posObj[1]] = {
                            index: posObj[1],
                            children: []
                        } ;
                    }
                    if (resultGrids[posObj[1]].children.indexOf(posObj[0])<0){
                        resultGrids[posObj[1]].children.push(posObj[0]) ;
                    }

                    // 生成意义显示索引
                    if (this.indexIdiomToCounter[item["idiom"]]==undefined) {
                        this.indexIdiomToCounter[item["idiom"]] = {
                            c: 0,
                            q: item["mean"]
                        } ;
                    }
                    this.indexIdiomToCounter[item["idiom"]].c++ ;
                    if (this.indexPosToIdiom[resultsKey]==undefined){
                        this.indexPosToIdiom[resultsKey] = [] ;
                    }
                    this.indexPosToIdiom[resultsKey].push(item["idiom"]) ;
                }
            }
        }
        let _resultGrids = [] ;
        for (let i in resultGrids) {
            resultGrids[i].children.sort();
            _resultGrids.push(resultGrids[i]) ;
        }

        let ids = [] ;

        // 渲染
        for (let i in data) {
            let item = data[i] ;
            let poses = item["pos"].split(";") ;

            for (let j=0,c=item["idiom"].length; j<c; j++) {
                let pos = poses.shift() ;
                let posObj = pos.split(",") ;
                let left = posObj[0]*span ;
                let top = posObj[1]*span ;
                let eleId = `item_${posObj[1]}_${posObj[0]}` ;

                if (ids.indexOf(eleId)>=0) {
                    continue ;
                }

                ids.push(eleId) ;

                let word = "?" ;
                let resultsKey = `${posObj[1]}-${posObj[0]}` ;

                if (results[resultsKey]==undefined) {
                    word = item["idiom"][j];
                }

                let ele = document.createElement("div") ;
                ele.setAttribute("class", "item");
                ele.setAttribute("style", `left: ${left}px; top: ${top}px`);
                ele.setAttribute("id", eleId);
                ele.append(word) ;

                document.getElementById("content").appendChild(ele) ;
            }
        }

        return {results: results, resultGrids: _resultGrids} ;
    }
} ;


