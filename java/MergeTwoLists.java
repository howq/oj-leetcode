public class MergeTwoLists {
    //Definition for singly-linked list.
    public static class ListNode {
        int val;
        ListNode next;
        ListNode(int x) { val = x; }
    }

    public static ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode head = new ListNode(0);
        ListNode temp = head;
        while (l1 != null && l2 != null) {
            if (l1.val < l2.val) {
                temp.next = l1;
                l1 = l1.next;
            } else {
                temp.next = l2;
                l2 = l2.next;
            }
            temp = temp.next;
        }
        temp.next = l1 != null ? l1 : l2;
        return head.next;
    }

    public static void main(String ... args){
        ListNode l1 = new ListNode(1);
        ListNode l11 = new ListNode(2);
        ListNode l12 = new ListNode(5);

        l1.next = l11;
        l11.next = l12;

        ListNode l2 = new ListNode(2);
        ListNode l21 = new ListNode(3);
        ListNode l22 = new ListNode(4);

        l2.next = l21;
        l21.next = l22;

        ListNode result = mergeTwoLists(l1, l2);


        while (result != null) {
            System.out.println(result.val);
            result = result.next;
        }

    }

}
