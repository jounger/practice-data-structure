function Node(data, next) {
    this.data = data
    this.next = next
}

function LinkedList(head) {
    this.head = head

    // Print out the list
    this.print = function() {
        let temp = this.head
        while(temp != null) {
            console.log(temp.data)
            temp = temp.next
        }
    }
    // Push in front of the list
    this.prepend = function(key) {
        const node = new Node(key)
        node.next = this.head
        this.head = next
    }
    // Append in end of the list
    this.append = function(key) {
        const node = new Node(key)
        node.next = null
        if(this.head == null) {
            this.head = node
            return
        }
        let temp = this.head
        while(temp != null) {
            if(temp.next == null) {
                temp.next = node
                return
            }
            temp = temp.next
        }
    }
    // Delete a node by given key
    this.delete = function(key) {
        let temp = this.head
        let prev = null
        if(temp != null && temp.data == key) {
            this.head = temp.next
            return
        }
        while(temp != null) {
            if(temp.data == key) {
                prev.next = temp.next
                return
            }
            prev = temp
            temp = temp.next
        }
    }
    // Delete a node by given index
    this.deleteIdx = function(idx) {
        let temp = this.head
        let prev = null
        if(temp != null && idx == 0) {
            this.head = temp.next
            return
        }
        for(let i = 0; i < idx && temp != null; i++) {
            if(i == idx) {
                prev.next = temp.next
                return
            }
            prev = temp
            temp = temp.next
        }
    }
    // Delete the list
    this.destroy = function(){
        this.head = null
    }
    // Count the length of list
    this.count = function() {
        let temp = this.head
        let count = 0
        while(temp != null) {
            count++
            temp = temp.next
        }
        return count
    }
    // Find an element in list and return boolean
    this.isExist = function(){
        // TODO:
    }
}

function main() {
    let list = new LinkedList(new Node(0))
    list.print()
}
main()