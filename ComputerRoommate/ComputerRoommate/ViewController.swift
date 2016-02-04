//
//  ViewController.swift
//  ComputerRoommate
//
//  Created by Dan Kurtz on 1/22/16.
//  Copyright © 2016 Dan Kurtz. All rights reserved.
//

import Cocoa

class ViewController: NSViewController {
    
    @IBAction func selectProfile(sender: NSButton) {
        let json = NSBundle.mainBundle().pathForResource(sender.title, ofType: "json")
        
        var resourcesParts = json!.characters.split("/").map(String.init)
        resourcesParts.removeAtIndex(resourcesParts.count-1)
        let resourcesPath = resourcesParts.joinWithSeparator("/")
        
        let path = NSBundle.mainBundle().pathForResource("roommate-cli", ofType: nil)
        let task = NSTask()
        task.launchPath = path
        task.arguments = ["-service="+json!, "-resources=/"+resourcesPath+"/"]
        task.launch()
        
        if(task.running == false) {
            let alert = NSAlert()
            alert.messageText = "Roommate service not running!"
            alert.informativeText = "The Roommate service was unable to start."
            alert.runModal()
        } else {
            NSApplication.sharedApplication().keyWindow?.close()
        }
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

