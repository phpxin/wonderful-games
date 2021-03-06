const SIZE = 8 ;
const SPAN = 30 ;

// todo 从网络获取，并且支持自动下一个谜语
let idioms = '[{"idiom":"苦思冥想","mean":"绞尽脑汁，深沉地思索。","pos":"3,6;4,6;5,6;6,6","vis":"1101"},{"idiom":"不辞劳苦","mean":"辞：推托。劳苦：劳累辛苦。不逃避劳累辛苦。形容人不怕吃苦，毅力强。","pos":"3,3;3,4;3,5;3,6","vis":"1101"},{"idiom":"迟疑不定","mean":"犹言迟疑不决。","pos":"1,3;2,3;3,3;4,3","vis":"1101"},{"idiom":"迟疑不决","mean":"形容拿不定主意。","pos":"1,3;1,4;1,5;1,6","vis":"1101"},{"idiom":"事不宜迟","mean":"事情要抓紧时机快做，不宜拖延。","pos":"1,0;1,1;1,2;1,3","vis":"1101"},{"idiom":"临事而惧","mean":"临：遭遇，碰到；惧：或惧。遇事谨慎戒惧。","pos":"0,0;1,0;2,0;3,0","vis":"0111"},{"idiom":"临深履薄","mean":"面临深渊，脚踩薄冰。比喻小心谨慎，惟恐有失。","pos":"0,0;0,1;0,2;0,3","vis":"1101"}]' ;
let idiomList = JSON.parse(idioms) ;
let tryCounter = 0 ;

createBackground(SIZE) ;

let {results,resultGrids} = createPlayground(SPAN, idiomList) ;
console.log("results : ", results) ;
console.log("resultGrids : ", JSON.stringify(resultGrids)) ;

renderResults(results) ;

let tmpResultRow = null ;
let questObj = quest() ;
let questEle = null ;

// `grid_${i}_${j}`
highlightGrid(questObj.i, questObj.j) ;

function fetchGame() {

}

// 点击结果
function onResultClicked(ele) {
    tryCounter++ ;
    let resultIndex = `${questObj.i}-${questObj.j}` ;
    let innerText = ele.childNodes[0].nodeValue ;
    if (results[resultIndex]!=innerText) {
        alert('错误，请重试') ;
        return
    }

    questEle.setAttribute("class", "bg-normal") ;
    ele.replaceWith("");
    document.getElementById(`item_${questObj.i}_${questObj.j}`).childNodes[0].replaceWith(innerText);

    // 下一个迷
    questObj = quest() ;
    if (questObj==null) {
        alert(`回答完成：共尝试了 ${tryCounter} 次`) ;
        window.location.reload() ;
    }
    highlightGrid(questObj.i, questObj.j) ;
}

// 高亮一个背景格子
function highlightGrid(i,j) {
    questEle = document.getElementById(`grid_${i}_${j}`);
    questEle.setAttribute("class", "bg-normal bg-highlight") ;
}

// 获取一个问题
function quest() {
    if (tmpResultRow==null || tmpResultRow.children.length<=0) {
        if (resultGrids.length<=0) {
            return null ;
        }
        tmpResultRow = resultGrids.shift() ;
    }

    let result = {} ;
    result.i = tmpResultRow.index ;
    result.j = tmpResultRow.children.shift() ;

    return result ;
}

// 结果集渲染
function renderResults(results) {
    let parentEle = document.getElementById("result") ;
    for (let i in results) {
        let ele = document.createElement("div") ;
        ele.setAttribute("class", "item") ;
        ele.setAttribute("onClick", "javascript:onResultClicked(this);") ;
        ele.append(results[i]) ;

        parentEle.appendChild(ele) ;
    }

    appendClean(parentEle)
}

// 辅助函数：添加清除样式
function appendClean(ele) {
    let clEle = document.createElement("div") ;
    clEle.setAttribute("class", "clean") ;
    ele.appendChild(clEle) ;
}

// 背景格子渲染
function createBackground(size) {
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
        appendClean(bgEle) ;
    }
}

// 主游戏逻辑界面渲染
function createPlayground(span, data) {
    let results = [] ;      // 结果集
    let resultGrids = [] ;  // 结果集格子数组

    for (let i in data) {
        let item = data[i] ;
        let poses = item["pos"].split(";") ;

        for (let j=0,c=item["idiom"].length; j<c; j++) {
            let pos = poses.shift() ;
            let posObj = pos.split(",") ;
            let resultsKey = `${posObj[1]}-${posObj[0]}` ;
            if (item["vis"][j]=="0") {
                results[resultsKey] = item["idiom"][j] ;

                if (resultGrids[posObj[1]]==undefined) {
                    resultGrids[posObj[1]] = {
                        index: posObj[1],
                        children: []
                    } ;
                }
                resultGrids[posObj[1]].children.push(posObj[0]) ;
            }
        }
    }
    let _resultGrids = [] ;
    for (let i in resultGrids) {
        resultGrids[i].children.sort();
        _resultGrids.push(resultGrids[i]) ;
    }

    let ids = [] ;

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

    return {results: results,resultGrids: _resultGrids} ;
}