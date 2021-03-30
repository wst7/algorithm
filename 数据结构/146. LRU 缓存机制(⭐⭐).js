/**
 * @param {number} capacity
 */
var LRUCache = function(capacity) {
    
  this.cap = capacity
  this.map = new Map()
  this.cache = new DoubleList()

};


  /** 
* @param {number} key
* @return {number}
*/
LRUCache.prototype.get = function(key) {
  if (!this.map.has(key)) {
      return -1
  }
  this.makeRecently(key)

  return this.map.get(key).val
};

/** 
* @param {number} key 
* @param {number} value
* @return {void}
*/
LRUCache.prototype.put = function(key, value) {
  // 已经存在key
  if (this.map.get(key)) {
      this.deleteKey(key)

      this.addRecently(key, value)
      return
  }

  if (this.cache.size === this.cap) {
      // 删除最久未使用的
      this.removeLeastRecently()
  }

  // 添加为最近使用的元素
  this.addRecently(key, value)
};





/**
* Your LRUCache object will be instantiated and called as such:
* var obj = new LRUCache(capacity)
* var param_1 = obj.get(key)
* obj.put(key,value)
*/

// 哈希链表：双向链表和哈希表的结合体
// 哈希表快速查找，链表快速插入和删除
class Node {
  next = null
  prev = null
  constructor(k, v) {
      this.key = k
      this.val = v
  }
}

class DoubleList {
  constructor() {
      this.head = new Node(0, 0)
      this.tail = new Node(0, 0)
      this.head.next = this.tail
      this.tail.prev = this.head
      this.size = 0
  }

  // 在链表尾部添加node
  addLast(node) {
      node.prev = this.tail.prev
      node.next = this.tail

      this.tail.prev.next = node
      this.tail.prev = node
      this.size ++
  }

  // 在链表中删除node（node一定存在）
  remove(node) {
      node.prev.next = node.next
      node.next.prev = node.prev

      this.size --
  }

  // 删除链表中第一个节点，并返回该节点
  removeFirst() {
      if (this.head === this.tail) {
          return
      }

      const first = this.head.next
      this.remove(first)

      return first
  }
}

// 将某个key提升为最近使用的
LRUCache.prototype.makeRecently = function(key) {
  const node = this.map.get(key)

  // 从链表中删除
  this.cache.remove(node)
  // 插入到尾部
  this.cache.addLast(node)
};

// 添加最近使用的
LRUCache.prototype.addRecently = function(key, val) {
  const node = new Node(key, val)

  // 插入到尾部
  this.cache.addLast(node)

  // 在map中添加key隐射
  this.map.set(key, node)
};

// 删除某一个key
LRUCache.prototype.deleteKey = function(key) {
  const node = this.map.get(key)

  // 从链表中删除
  this.cache.remove(node)

  // 重map中删除
  this.map.delete(key)
};

// 删除某一个key
LRUCache.prototype.removeLeastRecently = function() {
  const node = this.cache.removeFirst()

  const deleteKey = node.key

  this.map.delete(deleteKey)
};