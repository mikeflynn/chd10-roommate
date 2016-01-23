//
//  ViewController.swift
//  ComputerRoommate
//
//  Created by Dan Kurtz on 1/22/16.
//  Copyright © 2016 Dan Kurtz. All rights reserved.
//

import Cocoa

class ViewController: NSViewController {
    
    @IBAction func imageClicked(sender: AnyObject) {
        NSApplication.sharedApplication().keyWindow?.close()
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        self.view.wantsLayer = true
        self.view.layer!.backgroundColor = NSColor.whiteColor().CGColor
        // Do any additional setup after loading the view.
    }

    override var representedObject: AnyObject? {
        didSet {
        // Update the view, if already loaded.
        }
    }


}

